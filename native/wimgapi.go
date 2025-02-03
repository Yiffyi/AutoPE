package native

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	WIM_FLAG_RESERVED       = 0x00000001
	WIM_FLAG_VERIFY         = 0x00000002
	WIM_FLAG_INDEX          = 0x00000004
	WIM_FLAG_NO_APPLY       = 0x00000008
	WIM_FLAG_NO_DIRACL      = 0x00000010
	WIM_FLAG_NO_FILEACL     = 0x00000020
	WIM_FLAG_SHARE_WRITE    = 0x00000040
	WIM_FLAG_FILEINFO       = 0x00000080
	WIM_FLAG_NO_RP_FIX      = 0x00000100
	WIM_FLAG_MOUNT_READONLY = 0x00000200
	WIM_FLAG_MOUNT_FAST     = 0x00000400
	WIM_FLAG_MOUNT_LEGACY   = 0x00000800
	WIM_FLAG_APPLY_CI_EA    = 0x00001000
	WIM_FLAG_WIM_BOOT       = 0x00002000
	WIM_FLAG_APPLY_COMPACT  = 0x00004000
	WIM_FLAG_SUPPORT_EA     = 0x00008000
	// https://github.com/pbatard/rufus/commit/521034da991e0db4b0a77db9a6ea1c928a2eace8
	WIM_UNDOCUMENTED_BULLSHIT = 0x20000000
)

const (
	INVALID_CALLBACK_VALUE = 0xFFFFFFFF
)

const (
	WIM_COMPRESS_NONE = iota
	WIM_COMPRESS_XPRESS
	WIM_COMPRESS_LZX
)

const (
	WIM_MSG_SUCCESS     = 0
	WIM_MSG_DONE        = 0xFFFFFFF0
	WIM_MSG_SKIP_ERROR  = 0xFFFFFFFE
	WIM_MSG_ABORT_IMAGE = 0xFFFFFFFF
)

type WimMessageId uint32

const WM_APP = 32768
const (
	WIM_MSG WimMessageId = WM_APP + 0x1476 + iota
	WIM_MSG_TEXT
	WIM_MSG_PROGRESS
	WIM_MSG_PROCESS
	WIM_MSG_SCANNING
	WIM_MSG_SETRANGE
	WIM_MSG_SETPOS
	WIM_MSG_STEPIT
	WIM_MSG_COMPRESS
	WIM_MSG_ERROR
	WIM_MSG_ALIGNMENT
	WIM_MSG_RETRY
	WIM_MSG_SPLIT
	WIM_MSG_FILEINFO
	WIM_MSG_INFO
	WIM_MSG_WARNING
	WIM_MSG_CHK_PROCESS
	WIM_MSG_WARNING_OBJECTID
	WIM_MSG_STALE_MOUNT_DIR
	WIM_MSG_STALE_MOUNT_FILE
	WIM_MSG_MOUNT_CLEANUP_PROGRESS
	WIM_MSG_CLEANUP_SCANNING_DRIVE
	WIM_MSG_IMAGE_ALREADY_MOUNTED
	WIM_MSG_CLEANUP_UNMOUNTING_IMAGE
	WIM_MSG_QUERY_ABORT
	WIM_MSG_IO_RANGE_START_REQUEST_LOOP
	WIM_MSG_IO_RANGE_END_REQUEST_LOOP
	WIM_MSG_IO_RANGE_REQUEST
	WIM_MSG_IO_RANGE_RELEASE
	WIM_MSG_VERIFY_PROGRESS
	WIM_MSG_COPY_BUFFER
	WIM_MSG_METADATA_EXCLUDE
	WIM_MSG_GET_APPLY_ROOT
	WIM_MSG_MDPAD
	WIM_MSG_STEPNAME
	WIM_MSG_PERFILE_COMPRESS
	WIM_MSG_CHECK_CI_EA_PREREQUISITE_NOT_MET
	WIM_MSG_JOURNALING_ENABLED
)

type WimInfo struct {
	WimPath         [windows.MAX_PATH]uint16
	Guid            windows.GUID
	ImageCount      uint32
	CompressionType uint32
	PartNumber      uint16
	TotalParts      uint16
	BootIndex       uint32
	WimAttributes   uint32
	WimFlagsAndAttr uint32
}

type WIMMessageCallback func(dwMessageId WimMessageId, wParam uintptr, lParam uintptr, pvUserData uintptr) uint32

func WIMGetImageInformation(hImage syscall.Handle) (xml string, err error) {
	var bufImageInfo *uint16
	var cbImageInfo uint32
	err = WIMGetImageInformationR(hImage, &bufImageInfo, &cbImageInfo)
	if err != nil {
		return "", err
	}
	defer windows.LocalFree(windows.Handle(unsafe.Pointer(bufImageInfo)))

	// fmt.Println(err)
	xml = syscall.UTF16ToString(unsafe.Slice(bufImageInfo, cbImageInfo>>1))
	return

}

func WIMGetAttributes(hImage syscall.Handle) (wimInfo WimInfo, err error) {
	err = WIMGetAttributesR(hImage, &wimInfo, uint32(unsafe.Sizeof(wimInfo)))
	return
}
