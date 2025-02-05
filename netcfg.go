package autope

import (
	"errors"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"
	"golang.org/x/sys/windows/registry"
)

type NetworkConfig struct {
	DeviceInstanceID string

	EnableDHCP     bool
	IPAddress      []string
	SubnetMask     []string
	DefaultGateway []string
	NameServer     []string
}

func (nc *NetworkConfig) _fillWithInterfaceKeyDhcp(hInterfaceKey registry.Key) {
	var err error
	var dnsServers, ipAddrs, subnetMasks, defGateways []string

	strIPAddrs, _, err := hInterfaceKey.GetStringValue("DhcpIPAddress")
	if err != nil {
		ipAddrs = nil
	} else {
		ipAddrs = []string{strIPAddrs}
	}

	strSubnetMasks, _, err := hInterfaceKey.GetStringValue("DhcpSubnetMask")
	if err != nil {
		subnetMasks = nil
	} else {
		subnetMasks = []string{strSubnetMasks}
	}

	defGateways, _, err = hInterfaceKey.GetStringsValue("DhcpDefaultGateway")
	if err != nil {
		defGateways = nil
	}

	strDNSServers, _, err := hInterfaceKey.GetStringValue("DhcpNameServer")
	if err != nil {
		// strDNSServers =
		dnsServers = nil
	} else {
		if len(strDNSServers) > 0 { // avoid []string{""}
			dnsServers = strings.Split(strDNSServers, " ")
		} else {
			dnsServers = []string{}
		}
	}

	nc.EnableDHCP = true
	nc.IPAddress = ipAddrs
	nc.SubnetMask = subnetMasks
	nc.DefaultGateway = defGateways
	nc.NameServer = dnsServers
}

func (nc *NetworkConfig) _fillWithInterfaceKeyStatic(hInterfaceKey registry.Key) {
	var err error
	var dnsServers, ipAddrs, subnetMasks, defGateways []string

	ipAddrs, _, err = hInterfaceKey.GetStringsValue("IPAddress")
	if err != nil {
		ipAddrs = nil
	}

	subnetMasks, _, err = hInterfaceKey.GetStringsValue("SubnetMask")
	if err != nil {
		subnetMasks = nil
	}

	defGateways, _, err = hInterfaceKey.GetStringsValue("DefaultGateway")
	if err != nil {
		defGateways = nil
	}

	strDNSServers, _, err := hInterfaceKey.GetStringValue("NameServer")
	if err != nil {
		// strDNSServers =
		dnsServers = nil
	} else {
		if len(strDNSServers) > 0 { // avoid []string{""}
			dnsServers = strings.Split(strDNSServers, ",")
		} else {
			dnsServers = []string{}
		}
	}

	nc.EnableDHCP = false
	nc.IPAddress = ipAddrs
	nc.SubnetMask = subnetMasks
	nc.DefaultGateway = defGateways
	nc.NameServer = dnsServers
}

func (nc *NetworkConfig) fillWithInterfaceKey(hInterfaceKey registry.Key) error {
	enableDHCP, _, err := hInterfaceKey.GetIntegerValue("EnableDHCP")
	if err != nil {
		enableDHCP = 0
	}

	if enableDHCP > 0 {
		nc._fillWithInterfaceKeyDhcp(hInterfaceKey)
	} else {
		nc._fillWithInterfaceKeyStatic(hInterfaceKey)
	}

	return nil
}

func parseControlClassSubKey(hControlClassKey registry.Key, keyName string) (devId, netCfgId string, err error) {
	var digitCheck = regexp.MustCompile(`^[0-9]{4}$`)
	if !digitCheck.MatchString(keyName) {
		return "", "", errors.New("unwanted key name")
	}

	k, err := registry.OpenKey(hControlClassKey, keyName, registry.QUERY_VALUE)
	if err != nil {
		log.Error().Err(err).Str("keyName", keyName).Msg("failed to open key under Control\\Class")
		return "", "", err
		// return err
	}
	defer k.Close()

	devId, _, err = k.GetStringValue("DeviceInstanceID")
	if err != nil {
		return "", "", err
	}

	if strings.HasPrefix(devId, "SWD") {
		log.Debug().Str("DeviceInstanceID", devId).Msg("skipped SWD adapter")
		return "", "", errors.New("skipped SWD adapter")
	}

	if strings.HasPrefix(devId, "BTH") {
		log.Debug().Str("DeviceInstanceID", devId).Msg("skipped Bluetooth adapter")
		return "", "", errors.New("skipped Bluetooth adapter")
	}

	if strings.HasPrefix(devId, "ROOT\\KDNIC") {
		log.Debug().Str("DeviceInstanceID", devId).Msg("skipped KDNIC adapter")
		return "", "", errors.New("skipped KDNIC adapter")
	}

	netCfgId, _, err = k.GetStringValue("NetCfgInstanceId")
	if err != nil {
		return "", "", err
	}

	return
}

func GetNetworkConfigFromRegistry(controlSetPath string) (cfgs []*NetworkConfig, err error) {
	hKey, err := registry.OpenKey(registry.LOCAL_MACHINE, filepath.Join(controlSetPath, `Control\Class\{4d36e972-e325-11ce-bfc1-08002be10318}`), registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS)
	if err != nil {
		log.Error().Str("controlSet", controlSetPath).Err(err).Msg("failed to open key Control\\Class\\{4d36e972-e325-11ce-bfc1-08002be10318}")
		return nil, err
	}
	defer hKey.Close()

	hTcpipKey, err := registry.OpenKey(registry.LOCAL_MACHINE, filepath.Join(controlSetPath, `Services\Tcpip\Parameters`), registry.QUERY_VALUE)
	if err != nil {
		log.Error().Str("controlSet", controlSetPath).Err(err).Msg("failed to open key Services\\Tcpip\\Parameters")
		return
	}
	defer hTcpipKey.Close()

	subKeyNames, err := hKey.ReadSubKeyNames(0)
	if err != nil {
		log.Error().Err(err).Msg("hKey.ReadSubKeyNames")
		return nil, err
	}

	for _, keyName := range subKeyNames {
		devId, netCfgId, err := parseControlClassSubKey(hKey, keyName)
		if err != nil {
			continue
		}

		hInterfaceKey, err := registry.OpenKey(hTcpipKey, filepath.Join("Interfaces", netCfgId), registry.QUERY_VALUE)
		if err != nil {
			log.Error().Err(err).Str("NetCfgInstanceId", netCfgId).Msg("failed to open key under Tcpip\\Parameters\\Interfaces")
			continue
		}
		defer hInterfaceKey.Close()

		nc := &NetworkConfig{
			DeviceInstanceID: devId,
		}

		err = nc.fillWithInterfaceKey(hInterfaceKey)
		if err != nil {
			continue
		}

		log.Info().
			Interface("nc", nc).
			Msg("pickup netCfg")
		cfgs = append(cfgs, nc)
		// RegWrite($sPath, "IPAddress", "REG_MULTI_SZ", _ArrayToString($aIPAddr, @LF))
		// RegWrite($sPath, "SubnetMask", "REG_MULTI_SZ", _ArrayToString($aSubnet, @LF))
		// RegWrite($sPath, "DefaultGateway", "REG_MULTI_SZ", _ArrayToString($aDefGateway, @LF))
		// RegWrite($sPath, "NameServer", "REG_SZ", _ArrayToString($aDNS, ','))
	}
	return cfgs, nil
}
