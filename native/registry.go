package native

import (
	"syscall"

	"github.com/rs/zerolog/log"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

func RegLoadKey(hKey registry.Key, strSubKey string, strFile string) error {
	lpSubKey, err := syscall.UTF16PtrFromString(strSubKey)
	if err != nil {
		return err
	}

	lpFile, err := syscall.UTF16PtrFromString(strFile)
	if err != nil {
		return err
	}

	_, err = regLoadKey(syscall.Handle(hKey), lpSubKey, lpFile)
	if err != nil {
		return err
	}

	// if lstatus != uint32(windows.ERROR_SUCCESS) {
	// 	buf := make([]uint16, 256)
	// 	_, err := windows.FormatMessage(windows.FORMAT_MESSAGE_FROM_SYSTEM, 0, lstatus, 0, buf, nil)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return errors.New(syscall.UTF16ToString(buf))
	// }

	return nil
}

func RegUnLoadKey(hKey registry.Key, strSubKey string) error {
	lpSubKey, err := syscall.UTF16PtrFromString(strSubKey)
	if err != nil {
		return err
	}

	_, err = regUnLoadKey(syscall.Handle(hKey), lpSubKey)
	if err != nil {
		return err
	}

	// if lstatus != uint32(windows.ERROR_SUCCESS) {
	// 	buf := make([]uint16, 256)
	// 	_, err := windows.FormatMessage(windows.FORMAT_MESSAGE_FROM_SYSTEM, 0, lstatus, 0, buf, nil)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return errors.New(syscall.UTF16ToString(buf))
	// }

	return nil
}
func LoadHive(hivePath, subKeyName string) (err error) {

	var hToken windows.Token
	err = windows.OpenProcessToken(windows.CurrentProcess(), windows.TOKEN_ADJUST_PRIVILEGES, &hToken)
	if err != nil {
		log.Error().Err(err).Msg("OpenProcessToken failed")
		return err
	}

	err = SetPrivileges(hToken, []string{"SeBackupPrivilege", "SeRestorePrivilege"})
	if err != nil {
		log.Error().Err(err).Msg("SetPrivileges failed")
		return err
	}

	// err = native.RegLoadKey(syscall.HKEY_LOCAL_MACHINE, "OfflineWindows", filepath.Join(offlineDriveLetter, `\Windows\System32\config\SYSTEM`))
	err = RegLoadKey(syscall.HKEY_LOCAL_MACHINE, subKeyName, hivePath)
	if err != nil {
		log.Error().Err(err).Msg("RegLoadKey failed")
		return err
	}

	return nil
}
