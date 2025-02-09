package native

import (
	"unicode/utf16"
	"unsafe"
)

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

//sys regLoadKey(hKey syscall.Handle, lpSubKey *uint16, lpFile *uint16) (lstatus uint32, err error) [failretval!=windows.NO_ERROR] = Advapi32.RegLoadKeyW
//sys regUnLoadKey(hKey syscall.Handle, lpSubKey *uint16) (lstatus uint32, err error) [failretval!=windows.NO_ERROR] = Advapi32.RegUnLoadKeyW

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
