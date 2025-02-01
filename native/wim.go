package native

import (
	"os"
	"runtime"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type WIMMessageChannelList struct {
	Process  chan string
	Progress chan int
	ETA      chan uint64

	Others chan WimMessageId
	Quit   chan error
}

func WIMMessageToChan(dwMessageId WimMessageId, wParam *byte, lParam *byte, pvUserData *byte) uintptr {
	cs := (*WIMMessageChannelList)(unsafe.Pointer(pvUserData))

	switch dwMessageId {
	case WIM_MSG_PROCESS:
		pszFullPath := UTF16PtrToString((*uint16)(unsafe.Pointer(wParam)))
		var pfProcessFile *bool = (*bool)(unsafe.Pointer(lParam))
		*pfProcessFile = true

		cs.Process <- pszFullPath
		// fmt.Println(pszFullPath)
	case WIM_MSG_PROGRESS:
		/*
			wParam = (UINT) dwPercent;
			lParam = (UINT) dwTicksRemaining;
		*/
		dwPercent := int(uintptr(unsafe.Pointer(wParam)))
		dwTicksRemaining := uint64(uintptr(unsafe.Pointer(lParam)))

		cs.Progress <- dwPercent
		cs.ETA <- dwTicksRemaining
	default:
		cs.Others <- dwMessageId
		// fmt.Println("Received WIM Message:", dwMessageId)
	}
	return WIM_MSG_SUCCESS
}

var myWIMMessageCallback = syscall.NewCallback(WIMMessageToChan)

func WIMApplyImageByPath(wimPath string, imgIndex uint32, dst string, channels *WIMMessageChannelList) (err error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	pszWimPath, err := syscall.UTF16PtrFromString(wimPath)
	// fmt.Println(err)
	hWim, err := WIMCreateFile(pszWimPath, windows.GENERIC_READ, windows.OPEN_EXISTING, WIM_UNDOCUMENTED_BULLSHIT|WIM_FLAG_VERIFY, WIM_COMPRESS_NONE, nil)
	// fmt.Println(hWim, err)

	// wimInfo, err := WIMGetAttributes(hWim)
	// fmt.Println("wimInfo:", wimInfo, err)
	_, err = WIMRegisterMessageCallback(hWim, myWIMMessageCallback, uintptr(unsafe.Pointer(channels)))
	// fmt.Println("WIMRegisterMessageCallback:", idx, err)

	dname, err := os.MkdirTemp("", "autope_wimgapi_temp")
	// fmt.Println("os.MkdirTemp:", dname, err)
	pszTempPath, _ := syscall.UTF16PtrFromString(dname)
	err = WIMSetTemporaryPath(hWim, pszTempPath)
	// fmt.Println("WIMSetTemporaryPath:", err)

	hImage, err := WIMLoadImage(hWim, imgIndex)
	// fmt.Println("WIMLoadImage:", hImage, err)

	// xmlImageInfo, err := WIMGetImageInformation(hImage)
	// fmt.Println("WIMGetImageInformation:", xmlImageInfo, err)

	pszPath, _ := syscall.UTF16PtrFromString(dst)

	err = WIMApplyImage(hImage, pszPath, WIM_FLAG_NO_APPLY)
	// fmt.Println("WIMApplyImage:", err)

	channels.Quit <- err
	return
}
