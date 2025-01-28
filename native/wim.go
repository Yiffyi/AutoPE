package native

import (
	"fmt"
	"os"
	"runtime"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func WIMMessageToChan(dwMessageId WimMessageId, wParam uintptr, lParam uintptr, pvUserData uintptr) uintptr {
	fmt.Println("Received WIM Message:", dwMessageId)
	switch dwMessageId {
	case WIM_MSG_PROCESS:
		pszFullPath := UTF16PtrToString((*uint16)(unsafe.Pointer(wParam)))
		var pfProcessFile *bool = (*bool)(unsafe.Pointer(lParam))
		*pfProcessFile = true
		fmt.Println(pszFullPath)
	case WIM_MSG_PROGRESS:

	}
	return WIM_MSG_SUCCESS
}

var myWIMMessageCallback = syscall.NewCallback(WIMMessageToChan)

func WIMApplyImageByPath(wimPath string, imgIndex uint32, dst string) (err error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	pszWimPath, err := syscall.UTF16PtrFromString(wimPath)
	fmt.Println(err)
	hWim, err := WIMCreateFile(pszWimPath, windows.GENERIC_READ, windows.OPEN_EXISTING, WIM_UNDOCUMENTED_BULLSHIT|WIM_FLAG_VERIFY, WIM_COMPRESS_NONE, nil)
	fmt.Println(hWim, err)

	wimInfo, err := WIMGetAttributes(hWim)
	fmt.Println("wimInfo:", wimInfo, err)
	idx, err := WIMRegisterMessageCallback(hWim, myWIMMessageCallback, 0)
	fmt.Println("WIMRegisterMessageCallback:", idx, err)

	dname, err := os.MkdirTemp("", "autope_wimgapi_temp")
	fmt.Println("os.MkdirTemp:", dname, err)
	pszTempPath, _ := syscall.UTF16PtrFromString(dname)
	err = WIMSetTemporaryPath(hWim, pszTempPath)
	fmt.Println("WIMSetTemporaryPath:", err)

	hImage, err := WIMLoadImage(hWim, imgIndex)
	fmt.Println("WIMLoadImage:", hImage, err)

	xmlImageInfo, err := WIMGetImageInformation(hImage)
	fmt.Println("WIMGetImageInformation:", xmlImageInfo, err)

	pszPath, _ := syscall.UTF16PtrFromString(dst)

	err = WIMApplyImage(hImage, pszPath, WIM_FLAG_NO_APPLY)
	fmt.Println("WIMApplyImage:", err)
	return
}
