package native

import (
	"os"
	"runtime"
	"syscall"
	"unsafe"

	"github.com/rs/zerolog/log"
	"golang.org/x/sys/windows"
)

type WIMMessageContext struct {
	FileCount uint32
	FileIndex uint32

	Process  chan string
	Progress chan int
	ETA      chan uint32

	Others chan WimMessageId
	Quit   chan error

	Cancel bool
}

func WIMMessageToChan(dwMessageId WimMessageId, wParam *byte, lParam *byte, pvUserData *byte) uintptr {
	ctx := (*WIMMessageContext)(unsafe.Pointer(pvUserData))

	switch dwMessageId {
	case WIM_MSG_SETRANGE:
		dwFileCount := uint32(uintptr(unsafe.Pointer(lParam)))
		ctx.FileCount = dwFileCount
	case WIM_MSG_SETPOS:
		dwFileCount := uint32(uintptr(unsafe.Pointer(lParam)))
		ctx.FileIndex = dwFileCount
	case WIM_MSG_QUERY_ABORT:
		if ctx.Cancel {
			return WIM_MSG_ABORT_IMAGE
		} else {
			return WIM_MSG_SUCCESS
		}

	case WIM_MSG_PROCESS:
		pszFullPath := UTF16PtrToString((*uint16)(unsafe.Pointer(wParam)))
		var pfProcessFile *bool = (*bool)(unsafe.Pointer(lParam))
		*pfProcessFile = true

		ctx.Process <- pszFullPath
		// fmt.Println(pszFullPath)
		if ctx.Cancel {
			return WIM_MSG_ABORT_IMAGE
		} else {
			return WIM_MSG_SUCCESS
		}
	case WIM_MSG_PROGRESS:
		/*
			wParam = (UINT) dwPercent;
			lParam = (UINT) dwTicksRemaining;
		*/
		dwPercent := int(uintptr(unsafe.Pointer(wParam)))
		dwTicksRemaining := uint32(uintptr(unsafe.Pointer(lParam)))

		ctx.Progress <- dwPercent
		ctx.ETA <- dwTicksRemaining
	default:
		ctx.Others <- dwMessageId
		// fmt.Println("Received WIM Message:", dwMessageId)
	}
	return WIM_MSG_SUCCESS
}

var myWIMMessageCallback = syscall.NewCallback(WIMMessageToChan)

func WIMApplyImageByPath(wimPath string, imgIndex uint32, dst string, ctx *WIMMessageContext) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	var err error
	defer func() {
		ctx.Quit <- err
	}()

	pszWimPath, _ := syscall.UTF16PtrFromString(wimPath)

	// fmt.Println(err)
	hWim, err := WIMCreateFile(pszWimPath, windows.GENERIC_READ, windows.OPEN_EXISTING, WIM_UNDOCUMENTED_BULLSHIT|WIM_FLAG_VERIFY, WIM_COMPRESS_NONE, nil)
	if err != nil {
		log.Error().Err(err).Str("wimPath", wimPath).Msg("WIMCreateFile failed")
		return
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
	idx, err := WIMRegisterMessageCallback(hWim, myWIMMessageCallback, uintptr(unsafe.Pointer(ctx)))
	if err != nil {
		log.Error().Err(err).Msg("WIMRegisterMessageCallback failed")
		return
	}
	log.Debug().Uint32("idx", idx).Msg("registered myWIMMessageCallback")
	// fmt.Println("WIMRegisterMessageCallback:", idx, err)

	dname, err := os.MkdirTemp("", "autope_wimgapi_temp")
	if err != nil {
		log.Error().Err(err).Msg("os.MkdirTemp failed")
		return
	}
	// fmt.Println("os.MkdirTemp:", dname, err)
	log.Debug().Str("dir", dname).Msg("created temp dir")

	pszTempPath, _ := syscall.UTF16PtrFromString(dname)
	err = WIMSetTemporaryPath(hWim, pszTempPath)
	if err != nil {
		log.Error().Err(err).Msg("WIMSetTemporaryPath failed")
		return
	}
	// fmt.Println("WIMSetTemporaryPath:", err)

	hImage, err := WIMLoadImage(hWim, imgIndex)
	if err != nil {
		log.Error().Err(err).Msg("WIMLoadImage failed")
		return
	}
	// fmt.Println("WIMLoadImage:", hImage, err)

	// xmlImageInfo, err := WIMGetImageInformation(hImage)
	// fmt.Println("WIMGetImageInformation:", xmlImageInfo, err)

	pszPath, _ := syscall.UTF16PtrFromString(dst)

	err = WIMApplyImage(hImage, pszPath, 0)
	if err != nil {
		log.Error().Err(err).Msg("WIMApplyImage failed")
		return
	}
	// fmt.Println("WIMApplyImage:", err)

	return
}
