package autope

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rs/zerolog/log"
	"github.com/yiffyi/autope/native"
	"github.com/yiffyi/autope/tui"
	"github.com/yiffyi/autope/wmi"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

type PlaybookPEStage struct {
	PickupNetCfg    bool `comment:"提取网络配置"`
	WaitNetwork     bool `comment:"等待网络"`
	DisableFirewall bool `comment:"关防火墙"`

	DiskpartScript string `toml:",multiline" comment:"分区脚本"`
	BootVolume     string `comment:"启动分区"`
	SystemVolume   string `comment:"系统分区"`
	DataVolume     string `comment:"数据分区"`

	FormatSystem bool `comment:"格式化系统分区"`
	FormatData   bool `comment:"格式化数据分区"`
	FormatBoot   bool `comment:"格式化启动分区"`

	FixMBR     bool `comment:"修复MBR"`
	FixBootMgr bool `comment:"修复引导"`

	ImagePath  string `comment:"镜像路径"`
	ImageIndex int    `comment:"镜像编号"`
}

type Playbook struct {
	PEStage PlaybookPEStage `comment:"PE阶段"`
}

func formatVolumeWithWMI(svc *wmi.SWbemServices, chk func(string, *wmi.SWbemObject) (fs, label string)) (err error) {
	querySet, err := svc.InstancesOf("Win32_Volume")
	if err != nil {
		return err
	}

	volumes, err := querySet.ToSlice()
	if err != nil {
		return err
	}
	log.Info().Int("count", len(volumes)).Msg("found volumes")

	for _, v := range volumes {
		text, err := v.GetObjectText()
		if err == nil {
			log.Debug().Str("obj", text).Msg("looking at volume")
		}
		driveLetter := v.PropertyMustGetValue("DriveLetter").(string)

		// var fs, label string

		fs, label := chk(driveLetter, v)
		if len(fs) == 0 {
			continue
		}

		// ret, err := v.ExecMethod("Format", fs, bool(true), uint32(0), label, bool(false))
		inParam, err := wmi.GetMethodInParam(v, "Format")
		if err != nil {
			return err
		}

		err = inParam.PropertyPutValue("FileSystem", fs)
		if err != nil {
			return err
		}

		err = inParam.PropertyPutValue("Label", label)
		if err != nil {
			return err
		}

		err = inParam.PropertyPutValue("QuickFormat", true)
		if err != nil {
			return err
		}

		log.Debug().Str("inParam", inParam.String()).Msg("inParam")

		ret, err := v.ExecMethod_("Format", inParam)
		if err != nil {
			log.Error().Err(err).Str("DriveLetter", driveLetter).Msg("failed when formatting")
			return err
		}
		log.Info().Str("DriveLetter", driveLetter).Str("FileSystem", fs).Str("Label", label).Str("ReturnValue", ret.String()).Msg("formatted")
	}
	return errors.New("could not found drive letter")
}

func SearchOfflineWindows(svc *wmi.SWbemServices) (driveLetter string, err error) {
	querySet, err := svc.InstancesOf("Win32_Volume")
	if err != nil {
		return "", err
	}

	volumes, err := querySet.ToSlice()
	if err != nil {
		return "", err
	}
	log.Info().Int("count", len(volumes)).Msg("found volumes")

	systemDrive := os.Getenv("SystemDrive")
	if len(systemDrive) == 0 {
		systemDrive = os.Getenv("SystemRoot")[:2]
	}

	for _, v := range volumes {
		text, err := v.GetObjectText()
		if err == nil {
			log.Debug().Str("obj", text).Msg("looking at volume")
		}
		driveLetter = v.PropertyMustGetValue("DriveLetter").(string)

		if driveLetter == systemDrive {
			continue
		}

		stat, err := os.Stat(filepath.Join(driveLetter, `Windows\System32\config`))
		if err != nil || !stat.IsDir() {
			continue
		}

		return driveLetter, nil
	}

	return "", errors.New("no valid offline Windows")
}

func PickupNetCfg(offlineDriveLetter string) (err error) {

	var hToken windows.Token
	err = windows.OpenProcessToken(windows.CurrentProcess(), windows.TOKEN_ADJUST_PRIVILEGES, &hToken)
	if err != nil {
		log.Error().Err(err).Msg("OpenProcessToken failed")
		return err
	}

	err = native.SetPrivileges(hToken, []string{"SeBackupPrivilege", "SeRestorePrivilege"})
	if err != nil {
		log.Error().Err(err).Msg("SetPrivileges failed")
		return err
	}

	err = native.RegLoadKey(syscall.HKEY_LOCAL_MACHINE, "OfflineWindows", filepath.Join(offlineDriveLetter, `\Windows\System32\config\SYSTEM`))
	if err != nil {
		log.Error().Err(err).Msg("RegLoadKey failed")
		return err
	}

	hKey, err := registry.OpenKey(registry.LOCAL_MACHINE, filepath.Join(`OfflineWindows`, `ControlSet001\Control\Class\{4d36e972-e325-11ce-bfc1-08002be10318}`), registry.QUERY_VALUE)
	if err != nil {
		log.Error().Err(err).Msg("registry.OpenKey")
		return err
	}

	subKeyNames, err := hKey.ReadSubKeyNames(256)
	if err != nil {
		log.Error().Err(err).Msg("hKey.ReadSubKeyNames")
		return err
	}

	for _, keyName := range subKeyNames {
		k, err := registry.OpenKey(hKey, keyName, registry.QUERY_VALUE)
		if err != nil {
			log.Error().Err(err).Str("keyName", keyName).Msg("failed to open key under Control\\Class")
			continue
			// return err
		}

		devId, _, err := k.GetStringValue("DeviceInstanceID")
		if err != nil {
			continue
		}

		if strings.HasPrefix(devId, "SWD") {
			log.Debug().Str("DeviceInstanceID", devId).Msg("skipped SWD adapter")
			continue
		}

		netCfgId, _, err := k.GetStringValue("NetCfgInstanceId")
		if err != nil {
			continue
		}

		kTcpip, err := registry.OpenKey(registry.LOCAL_MACHINE, filepath.Join(`OfflineWindows`, `ControlSet001\Services\Tcpip\Parameters\Interfaces\`, netCfgId), registry.QUERY_VALUE)
		if err != nil {
			log.Error().Err(err).Str("NetCfgInstanceId", netCfgId).Msg("failed to open key under Tcpip\\Parameters\\Interfaces")
			continue
		}

		enableDHCP, _, err := kTcpip.GetIntegerValue("EnableDHCP")
		if err != nil {
			enableDHCP = 0
		}

		ipAddrs, _, err := kTcpip.GetStringsValue("IPAddress")
		if err != nil {
			ipAddrs = nil
		}

		subnetMasks, _, err := kTcpip.GetStringsValue("SubnetMask")
		if err != nil {
			subnetMasks = nil
		}

		defGateway, _, err := kTcpip.GetStringsValue("DefaultGateway")
		if err != nil {
			defGateway = nil
		}

		strDNSServers, _, err := kTcpip.GetStringValue("NameServer")
		dnsServers := strings.Split(strDNSServers, ",")
		if err != nil {
			// strDNSServers =
			dnsServers = nil
		}
		log.Info().
			Bool("EnableDHCP", enableDHCP > 0).
			Strs("IPAddress", ipAddrs).
			Strs("SubnetMask", subnetMasks).
			Strs("DefaultGateway", defGateway).
			Strs("NameServer", dnsServers).
			Str("DeviceInstanceID", devId).
			Str("NetCfgInstanceId", netCfgId).
			Msg("pickup netCfg")
		// RegWrite($sPath, "IPAddress", "REG_MULTI_SZ", _ArrayToString($aIPAddr, @LF))
		// RegWrite($sPath, "SubnetMask", "REG_MULTI_SZ", _ArrayToString($aSubnet, @LF))
		// RegWrite($sPath, "DefaultGateway", "REG_MULTI_SZ", _ArrayToString($aDefGateway, @LF))
		// RegWrite($sPath, "NameServer", "REG_SZ", _ArrayToString($aDNS, ','))
	}
	return nil
}

func (p *PlaybookPEStage) Run() (err error) {

	if is, err := IsMiniNT(); !is {
		return fmt.Errorf("this is not suppose run here: %w", err)
	}

	var cmd *exec.Cmd

	if p.WaitNetwork {
		cmd = exec.Command("wpeutil", "WaitForNetwork")
		err = cmd.Run()
		if err != nil {
			return
		}
	}

	if p.DisableFirewall {
		cmd = exec.Command("wpeutil", "DisableFirewall")
		err = cmd.Run()
		if err != nil {
			return
		}
	}

	if len(p.DiskpartScript) > 0 {
		var fd *os.File
		fd, err = os.CreateTemp("", "autope_diskpart")
		if err != nil {
			return err
		}

		scriptPath := fd.Name()
		fd.WriteString(p.DiskpartScript)
		fd.Close()

		cmd = exec.Command("diskpart", "/s", scriptPath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			return
		}
	}

	wmi.CoInitialize()
	defer wmi.CoUninitialize()

	locator, err := wmi.NewSWbemLocator()
	if err != nil {
		log.Error().Err(err).Msg("could not initialize WMI service locator")
		return err
		// return nil, err
	}

	svc, err := locator.ConnectServerDefault()
	if err != nil {
		log.Error().Err(err).Msg("could not connect to WMI service")
		return err
	}

	formatVolumeWithWMI(svc, func(driveLetter string, v *wmi.SWbemObject) (fs, label string) {
		/*
			in WMI, Win32_Volume,
			BootVolume means the volume contains Windows
			SystemVolume means the volume contains bootloader
		*/
		if p.FormatBoot && driveLetter == filepath.VolumeName(p.BootVolume) {
			fs = "FAT32"
			label = "BOOT"

			if !v.PropertyMustGetValue("SystemVolume").(bool) {
				log.Warn().Str("driveLetter", driveLetter).Msg("is not considered as a boot volume by WMI")
			}
		} else if p.FormatSystem && driveLetter == filepath.VolumeName(p.SystemVolume) {
			fs = "NTFS"
			label = "SYSTEM"

			if !v.PropertyMustGetValue("BootVolume").(bool) {
				log.Warn().Str("driveLetter", driveLetter).Msg("is not considered as a system volume by WMI")
			}
		} else {
			fs = ""
			label = ""
		}
		return
	})

	if len(p.ImagePath) > 0 {
		log.Info().
			Str("imagePath", p.ImagePath).
			Int("imageIndex", p.ImageIndex).
			Str("systemVolume", p.SystemVolume).
			Msg("apply WIM")

		ctx := &native.WIMMessageContext{
			Process:  make(chan string, 16),
			Progress: make(chan int),
			ETA:      make(chan uint32),

			Others: make(chan native.WimMessageId, 16),
			Quit:   make(chan error),
		}

		go native.WIMApplyImageByPath(p.ImagePath, uint32(p.ImageIndex), p.SystemVolume, ctx)

		p := tea.NewProgram(tui.CreateTUIAppltImage(ctx))
		SetupBubbleTeaLogger(p)
		m, err := p.Run()
		SetupDefaultLogger()
		if err != nil {
			return err
		}
		tui := m.(*tui.TUIApplyImage)
		if tui.Error != nil {
			return tui.Error
		}
	}

	// delay the format of DataVolume, so that we can leave something in DataVolume
	if p.FormatData {
		formatVolumeWithWMI(svc, func(driveLetter string, _ *wmi.SWbemObject) (fs, label string) {
			if driveLetter == filepath.VolumeName(p.DataVolume) {
				fs = "NTFS"
				label = "DATA"
			} else {
				fs = ""
				label = ""
			}
			return
		})
	}

	if p.FixMBR {
		cmd = exec.Command("bootsect", "/nt60", filepath.VolumeName(p.BootVolume), "/mbr", "/force")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			log.Error().Err(err).Msg("writing MBR & PBR")
		} else {
			log.Info().Msg("written MBR & PBR")
		}
	}

	if p.FixBootMgr {
		winDir := filepath.Join(p.SystemVolume, `\Windows`)
		if len(p.BootVolume) > 0 {
			cmd = exec.Command("bcdboot", winDir, "/s", p.BootVolume)
		} else {
			cmd = exec.Command("bcdboot", winDir)
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			log.Error().Err(err).Msg("configuring bootmgr")
		} else {
			log.Info().Str("winDir", winDir).Msg("configured bootmgr")
		}
	}
	return nil
}
