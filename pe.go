package autope

import "golang.org/x/sys/windows/registry"

func IsMiniNT() (bool, error) {
	hKey, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlset\Control\MiniNT`, registry.READ)
	if err != nil {
		return false, err
	}
	defer hKey.Close()

	// TODO: check details
	return true, nil
}
