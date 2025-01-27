package native

import (
	"syscall"
	"unicode/utf16"
	"unsafe"

	"golang.org/x/sys/windows"
)

/*
// WIMCreateFile:
//
#define WIM_GENERIC_READ            GENERIC_READ
#define WIM_GENERIC_WRITE           GENERIC_WRITE
#define WIM_GENERIC_MOUNT           GENERIC_EXECUTE

#define WIM_CREATE_NEW              CREATE_NEW
#define WIM_CREATE_ALWAYS           CREATE_ALWAYS
#define WIM_OPEN_EXISTING           OPEN_EXISTING
#define WIM_OPEN_ALWAYS             OPEN_ALWAYS

enum
{
    WIM_COMPRESS_NONE = 0,
    WIM_COMPRESS_XPRESS = 1,
    WIM_COMPRESS_LZX = 2,
    WIM_COMPRESS_LZMS = 3
};

enum
{
    WIM_CREATED_NEW = 0,
    WIM_OPENED_EXISTING
};

// WIMCreateFile, WIMCaptureImage, WIMApplyImage, WIMMountImageHandle flags:
//
#define WIM_FLAG_RESERVED                  0x00000001
#define WIM_FLAG_VERIFY                    0x00000002
#define WIM_FLAG_INDEX                     0x00000004
#define WIM_FLAG_NO_APPLY                  0x00000008
#define WIM_FLAG_NO_DIRACL                 0x00000010
#define WIM_FLAG_NO_FILEACL                0x00000020
#define WIM_FLAG_SHARE_WRITE               0x00000040
#define WIM_FLAG_FILEINFO                  0x00000080
#define WIM_FLAG_NO_RP_FIX                 0x00000100
#define WIM_FLAG_MOUNT_READONLY            0x00000200
#define WIM_FLAG_MOUNT_FAST                0x00000400
#define WIM_FLAG_MOUNT_LEGACY              0x00000800
#define WIM_FLAG_APPLY_CI_EA               0x00001000
#define WIM_FLAG_WIM_BOOT                  0x00002000
#define WIM_FLAG_APPLY_COMPACT             0x00004000
#define WIM_FLAG_SUPPORT_EA                0x00008000 // It can be used in mount also.

// WIMGetMountedImageList flags
//
#define WIM_MOUNT_FLAG_MOUNTED              0x00000001
#define WIM_MOUNT_FLAG_MOUNTING             0x00000002
#define WIM_MOUNT_FLAG_REMOUNTABLE          0x00000004
#define WIM_MOUNT_FLAG_INVALID              0x00000008
#define WIM_MOUNT_FLAG_NO_WIM               0x00000010
#define WIM_MOUNT_FLAG_NO_MOUNTDIR          0x00000020
#define WIM_MOUNT_FLAG_MOUNTDIR_REPLACED    0x00000040
#define WIM_MOUNT_FLAG_READWRITE            0x00000100

// WIMCommitImageHandle flags
//
#define WIM_COMMIT_FLAG_APPEND      0x00000001

// WIMSetReferenceFile
//
#define WIM_REFERENCE_APPEND        0x00010000
#define WIM_REFERENCE_REPLACE       0x00020000

// WIMExportImage
//
#define WIM_EXPORT_ALLOW_DUPLICATES   0x00000001
#define WIM_EXPORT_ONLY_RESOURCES     0x00000002
#define WIM_EXPORT_ONLY_METADATA      0x00000004
#define WIM_EXPORT_VERIFY_SOURCE      0x00000008
#define WIM_EXPORT_VERIFY_DESTINATION 0x00000010

// WIMRegisterMessageCallback:
//
#define INVALID_CALLBACK_VALUE      0xFFFFFFFF

// WIMCopyFile
//
#define WIM_COPY_FILE_RETRY         0x01000000

// WIMDeleteImageMounts
//
#define WIM_DELETE_MOUNTS_ALL       0x00000001

// WIMRegisterLogfile
//
#define WIM_LOGFILE_UTF8            0x00000001

// WIMMessageCallback Notifications:
//
enum
{
    WIM_MSG = WM_APP + 0x1476,
    WIM_MSG_TEXT,
    WIM_MSG_PROGRESS,
    WIM_MSG_PROCESS,
    WIM_MSG_SCANNING,
    WIM_MSG_SETRANGE,
    WIM_MSG_SETPOS,
    WIM_MSG_STEPIT,
    WIM_MSG_COMPRESS,
    WIM_MSG_ERROR,
    WIM_MSG_ALIGNMENT,
    WIM_MSG_RETRY,
    WIM_MSG_SPLIT,
    WIM_MSG_FILEINFO,
    WIM_MSG_INFO,
    WIM_MSG_WARNING,
    WIM_MSG_CHK_PROCESS,
    WIM_MSG_WARNING_OBJECTID,
    WIM_MSG_STALE_MOUNT_DIR,
    WIM_MSG_STALE_MOUNT_FILE,
    WIM_MSG_MOUNT_CLEANUP_PROGRESS,
    WIM_MSG_CLEANUP_SCANNING_DRIVE,
    WIM_MSG_IMAGE_ALREADY_MOUNTED,
    WIM_MSG_CLEANUP_UNMOUNTING_IMAGE,
    WIM_MSG_QUERY_ABORT,
    WIM_MSG_IO_RANGE_START_REQUEST_LOOP,
    WIM_MSG_IO_RANGE_END_REQUEST_LOOP,
    WIM_MSG_IO_RANGE_REQUEST,
    WIM_MSG_IO_RANGE_RELEASE,
    WIM_MSG_VERIFY_PROGRESS,
    WIM_MSG_COPY_BUFFER,
    WIM_MSG_METADATA_EXCLUDE,
    WIM_MSG_GET_APPLY_ROOT,
    WIM_MSG_MDPAD,
    WIM_MSG_STEPNAME,
    WIM_MSG_PERFILE_COMPRESS,
    WIM_MSG_CHECK_CI_EA_PREREQUISITE_NOT_MET,
    WIM_MSG_JOURNALING_ENABLED
};

//
// WIMMessageCallback Return codes:
//
#define WIM_MSG_SUCCESS          ERROR_SUCCESS
#define WIM_MSG_DONE             0xFFFFFFF0
#define WIM_MSG_SKIP_ERROR       0xFFFFFFFE
#define WIM_MSG_ABORT_IMAGE      0xFFFFFFFF

//
// WIM_INFO dwFlags values:
//
#define WIM_ATTRIBUTE_NORMAL        0x00000000
#define WIM_ATTRIBUTE_RESOURCE_ONLY 0x00000001
#define WIM_ATTRIBUTE_METADATA_ONLY 0x00000002
#define WIM_ATTRIBUTE_VERIFY_DATA   0x00000004
#define WIM_ATTRIBUTE_RP_FIX        0x00000008
#define WIM_ATTRIBUTE_SPANNED       0x00000010
#define WIM_ATTRIBUTE_READONLY      0x00000020

//
// The WIM_INFO structure used by WIMGetAttributes:
//
typedef struct _WIM_INFO
{
    WCHAR  WimPath[MAX_PATH];
    GUID   Guid;
    DWORD  ImageCount;
    DWORD  CompressionType;
    USHORT PartNumber;
    USHORT TotalParts;
    DWORD  BootIndex;
    DWORD  WimAttributes;
    DWORD  WimFlagsAndAttr;
} WIM_INFO, *PWIM_INFO, *LPWIM_INFO;

//
// The WIM_MOUNT_LIST structure used for getting the list of mounted images.
//
typedef struct _WIM_MOUNT_LIST
{
    WCHAR  WimPath[MAX_PATH];
    WCHAR  MountPath[MAX_PATH];
    DWORD  ImageIndex;
    BOOL   MountedForRW;
} WIM_MOUNT_LIST, *PWIM_MOUNT_LIST, *LPWIM_MOUNT_LIST,
  WIM_MOUNT_INFO_LEVEL0, *PWIM_MOUNT_INFO_LEVEL0, LPWIM_MOUNT_INFO_LEVEL0;

//
// Define new WIM_MOUNT_INFO_LEVEL1 structure with additional data...
//
typedef struct _WIM_MOUNT_INFO_LEVEL1
{
    WCHAR  WimPath[MAX_PATH];
    WCHAR  MountPath[MAX_PATH];
    DWORD  ImageIndex;
    DWORD  MountFlags;
} WIM_MOUNT_INFO_LEVEL1, *PWIM_MOUNT_INFO_LEVEL1, *LPWIM_MOUNT_INFO_LEVEL1;

typedef WIM_MOUNT_INFO_LEVEL1 WIM_MOUNT_INFO_LATEST, *PWIM_MOUNT_INFO_LATEST;

//
// Define enumeration for WIMGetMountedImageInfo to determine structure to use...
//
typedef enum _MOUNTED_IMAGE_INFO_LEVELS
{
    MountedImageInfoLevel0,
    MountedImageInfoLevel1,
    MountedImageInfoLevelInvalid
} MOUNTED_IMAGE_INFO_LEVELS;

//
// An abstract type implemented by the caller when using File I/O callbacks.
//
typedef VOID * PFILEIOCALLBACK_SESSION;

//
// The WIM_IO_RANGE_CALLBACK structure is used in conjunction with the
// FileIOCallbackReadFile callback and the WIM_MSG_IO_RANGE_REQUEST and
// WIM_MSG_IO_RANGE_RELEASE message callbacks.  A pointer to a
// WIM_IO_RANGE_REQUEST is passed in WPARAM to the callback for both messages.
//
typedef struct _WIM_IO_RANGE_CALLBACK
{
    //
    // The callback session that corresponds to the file that is being queried.
    //
    PFILEIOCALLBACK_SESSION pSession;

    // Filled in by WIMGAPI for both messages:
    LARGE_INTEGER Offset, Size;

    // Filled in by the callback for WIM_MSG_IO_RANGE_REQUEST (set to TRUE to
    // indicate data in the specified range is available, and FALSE to indicate
    // it is not yet available):
    BOOL Available;
} WIM_IO_RANGE_CALLBACK, *PWIM_IO_RANGE_CALLBACK;


//
// Abstract (opaque) type for WIM files used with
// WIMEnumImageFiles API
//
typedef VOID * PWIM_ENUM_FILE;


#if defined(__cplusplus)
typedef struct _WIM_FILE_FIND_DATA : public _WIN32_FIND_DATAW
{
#else
typedef struct _WIM_FILE_FIND_DATA
{
    WIN32_FIND_DATAW;
#endif

    BYTE bHash[20];
    PSECURITY_DESCRIPTOR pSecurityDescriptor;
    PWSTR *ppszAlternateStreamNames;
    PBYTE pbReparseData;
    DWORD cbReparseData;
    ULARGE_INTEGER uliResourceSize;
} WIM_FIND_DATA, *PWIM_FIND_DATA;

HANDLE
WINAPI
WIMCreateFile(
    PWSTR   pszWimPath,
    DWORD   dwDesiredAccess,
    DWORD   dwCreationDisposition,
    DWORD   dwFlagsAndAttributes,
    DWORD   dwCompressionType,
    PDWORD  pdwCreationResult
    );

BOOL
WINAPI
WIMCloseHandle(
    HANDLE hObject
    );

BOOL
WINAPI
WIMGetAttributes(
    HANDLE     hWim,
    PWIM_INFO  pWimInfo,
    DWORD      cbWimInfo
    );

DWORD
WINAPI
WIMGetImageCount(
    HANDLE hWim
    );

DWORD
WINAPI
WIMGetMessagecallbackCount(
      HANDLE hWim
    );

HANDLE
WIMLoadImage(
    HANDLE hWim,
    DWORD  dwImageIndex
    );

BOOL
WINAPI
WIMGetImageInformation(
     HANDLE  hImage,
     PVOID  *ppvImageInfo,
     PDWORD  pcbImageInfo
    );

BOOL
WINAPI
WIMApplyImage(
    HANDLE hImage,
    PCWSTR pszPath,
    DWORD  dwApplyFlags
    );

DWORD
CALLBACK
WIMMessageCallback(
    DWORD  dwMessageId,
    WPARAM wParam,
    LPARAM lParam,
    PVOID  pvUserData
    );

HANDLE
WIMLoadImage(
    HANDLE hWim,
    DWORD  dwImageIndex
    );

BOOL
WIMSetTemporaryPath(
    HANDLE  hWim,
    PWSTR   pszPath
    );
*/

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

//sys WIMCreateFile(pszWimPath *uint16, dwDesiredAccess uint32, dwCreationDisposition uint32, dwFlagsAndAttributes uint32, dwCompressionType uint32, pdwCreationResult *uint32) (handle syscall.Handle, err error) = wimgapi.WIMCreateFile
//sys WIMCloseHandle(hObject syscall.Handle) (err error) = wimgapi.WIMCloseHandle

//sys WIMGetImageCount(hWim syscall.Handle) (n uint32, err error) = wimgapi.WIMGetImageCount
//sys WIMGetMessagecallbackCount(hWim syscall.Handle) (n uint32, err error) = wimgapi.WIMGetMessagecallbackCount
//sys WIMRegisterMessageCallback(hWim syscall.Handle, fpMessageProc uintptr, pvUserData uintptr) (idx uint32, err error) [failretval==INVALID_CALLBACK_VALUE] = wimgapi.WIMRegisterMessageCallback
//sys WIMSetTemporaryPath(hWim syscall.Handle, pszPath *uint16) (err error) = wimgapi.WIMSetTemporaryPath
//sys WIMLoadImage(hWim syscall.Handle, dwImageIndex uint32) (hImage syscall.Handle, err error) = wimgapi.WIMLoadImage

//sys WIMGetImageInformationR(hImage syscall.Handle, ppvImageInfo **uint16, pcbImageInfo *uint32) (err error) = wimgapi.WIMGetImageInformation
//sys WIMGetAttributesR(hImage syscall.Handle, pWimInfo *WimInfo, cbWimInfo uint32) (err error) = wimgapi.WIMGetAttributes
//sys WIMApplyImage(hImage syscall.Handle, pszPath *uint16, dwApplyFlags uint32) (err error) = wimgapi.WIMApplyImage

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

// utf16PtrToString is like UTF16ToString, but takes *uint16
// as a parameter instead of []uint16.
func UTF16PtrToString(p *uint16) string {
	if p == nil {
		return ""
	}
	// Find NUL terminator.
	end := unsafe.Pointer(p)
	n := 0
	for *(*uint16)(end) != 0 {
		end = unsafe.Pointer(uintptr(end) + unsafe.Sizeof(*p))
		n++
	}
	// Turn *uint16 into []uint16.
	s := unsafe.Slice(p, n)
	// Decode []uint16 into string.
	return string(utf16.Decode(s))
}
