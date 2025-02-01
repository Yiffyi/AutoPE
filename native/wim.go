package native

import (
	"os"
	"runtime"
	"syscall"
	"unsafe"

	"github.com/rs/zerolog/log"
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
	defer func() {
		channels.Quit <- err
	}()

	pszWimPath, _ := syscall.UTF16PtrFromString(wimPath)

	// fmt.Println(err)
	hWim, err := WIMCreateFile(pszWimPath, windows.GENERIC_READ, windows.OPEN_EXISTING, WIM_UNDOCUMENTED_BULLSHIT|WIM_FLAG_VERIFY, WIM_COMPRESS_NONE, nil)
	if err != nil {
		log.Error().Err(err).Str("wimPath", wimPath).Msg("WIMCreateFile failed")
		return err
	}
	// fmt.Println(hWim, err)

	wimInfo, err := WIMGetAttributes(hWim)
	log.Debug().
		Str("path", syscall.UTF16ToString(wimInfo.WimPath[:])).
		Uint32("imageCount", wimInfo.ImageCount).
		Uint32("compressionType", wimInfo.CompressionType).
		Uint16("partNumber", wimInfo.PartNumber).
		Uint16("totalParts", wimInfo.TotalParts).
		Msg("WIMGetAttributes")

	// fmt.Println("wimInfo:", wimInfo, err)
	idx, err := WIMRegisterMessageCallback(hWim, myWIMMessageCallback, uintptr(unsafe.Pointer(channels)))
	if err != nil {
		log.Error().Err(err).Msg("WIMRegisterMessageCallback failed")
		return err
	}
	log.Debug().Uint32("idx", idx).Msg("registered myWIMMessageCallback")
	// fmt.Println("WIMRegisterMessageCallback:", idx, err)

	dname, err := os.MkdirTemp("", "autope_wimgapi_temp")
	if err != nil {
		log.Error().Err(err).Msg("os.MkdirTemp failed")
		return err
	}
	// fmt.Println("os.MkdirTemp:", dname, err)
	log.Debug().Str("dir", dname).Msg("created temp dir")

	pszTempPath, _ := syscall.UTF16PtrFromString(dname)
	err = WIMSetTemporaryPath(hWim, pszTempPath)
	if err != nil {
		log.Error().Err(err).Msg("WIMSetTemporaryPath failed")
		return err
	}
	// fmt.Println("WIMSetTemporaryPath:", err)

	hImage, err := WIMLoadImage(hWim, imgIndex)
	if err != nil {
		log.Error().Err(err).Msg("WIMLoadImage failed")
		return err
	}
	// fmt.Println("WIMLoadImage:", hImage, err)

	// xmlImageInfo, err := WIMGetImageInformation(hImage)
	// fmt.Println("WIMGetImageInformation:", xmlImageInfo, err)

	pszPath, _ := syscall.UTF16PtrFromString(dst)

	err = WIMApplyImage(hImage, pszPath, WIM_FLAG_NO_APPLY)
	if err != nil {
		log.Error().Err(err).Msg("WIMApplyImage failed")
		return err
	}
	// fmt.Println("WIMApplyImage:", err)

	return
}
