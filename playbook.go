package autope

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/yiffyi/autope/native"
)

type PlaybookPEStage struct {
	WaitNetwork     bool `comment:"等待网络"`
	DisableFirewall bool `comment:"关防火墙"`

	DiskpartScript string `toml:",multiline" comment:"分区脚本"`
	BootVolume     string `comment:"启动分区"`
	SystemVolume   string `comment:"系统分区"`
	DataVolume     string `comment:"数据分区"`

	FormatSystem bool `comment:"格式化系统分区"`
	FormatData   bool `comment:"格式化数据分区"`
	FormatBoot   bool `comment:"格式化启动分区"`

	ImagePath  string `comment:"镜像路径"`
	ImageIndex int    `comment:"镜像编号"`
}

type Playbook struct {
	PEStage PlaybookPEStage `comment:"PE阶段"`
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
		err = cmd.Run()
		if err != nil {
			return
		}
	}

	if len(p.ImagePath) > 0 {
		err = native.WIMApplyImageByPath(p.ImagePath, uint32(p.ImageIndex), p.SystemVolume)
		if err != nil {
			return
		}
	}
	return nil
}
