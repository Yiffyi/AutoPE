package native

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type TokenPrivileges12 struct {
	PrivilegeCount uint32
	Privileges     [12]windows.LUIDAndAttributes // Use AllPrivileges() for iterating.
}

var (
	modadvapi32 = windows.NewLazySystemDLL("advapi32.dll")

	procAdjustTokenPrivileges = modadvapi32.NewProc("AdjustTokenPrivileges")
)

func AdjustTokenPrivileges(token windows.Token, disableAllPrivileges bool, newstate *windows.Tokenprivileges, buflen uint32, prevstate *windows.Tokenprivileges, returnlen *uint32) (success bool, err error) {
	var _p0 uint32
	if disableAllPrivileges {
		_p0 = 1
	}
	r1, _, e1 := syscall.SyscallN(procAdjustTokenPrivileges.Addr(), uintptr(token), uintptr(_p0), uintptr(unsafe.Pointer(newstate)), uintptr(buflen), uintptr(unsafe.Pointer(prevstate)), uintptr(unsafe.Pointer(returnlen)))
	success = r1 != 0
	err = e1
	return
}

func SetPrivileges(hToken windows.Token, wantedPrivilegeNames []string) (err error) {
	var tp TokenPrivileges12
	if len(wantedPrivilegeNames) > len(tp.Privileges) {
		panic("This is more than we could handle")
	}

	var s *uint16
	tp.PrivilegeCount = uint32(len(wantedPrivilegeNames))

	for k, v := range wantedPrivilegeNames {
		s, err = syscall.UTF16PtrFromString(v)
		if err != nil {
			return
		}

		err = windows.LookupPrivilegeValue(nil, s, &tp.Privileges[k].Luid)
		if err != nil {
			return
		}

		tp.Privileges[k].Attributes = windows.SE_PRIVILEGE_ENABLED
	}

	success, err := AdjustTokenPrivileges(hToken, false, (*windows.Tokenprivileges)(unsafe.Pointer(&tp)), uint32(unsafe.Sizeof(tp)), nil, nil)
	if success && err != windows.ERROR_SUCCESS {
		return
	}
	return nil
}
