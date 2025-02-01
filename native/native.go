package native

import (
	"unicode/utf16"
	"unsafe"
)

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
