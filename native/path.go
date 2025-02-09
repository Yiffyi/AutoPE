package native

import (
	"syscall"
	"unicode/utf16"

	"golang.org/x/sys/windows"
)

func bitsToDrives(bitMap uint32) (drives []string) {
	availableDrives := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	for i := range availableDrives {
		if bitMap&1 == 1 {
			drives = append(drives, availableDrives[i]+":")
		}
		bitMap >>= 1
	}

	return
}

func utf16ToStrings(buf []uint16) (result []string) {
	var s string
	posStrBegin := 0
	for idx, v := range buf {
		if v == 0 {
			if idx == posStrBegin {
				break
			}
			s = string(utf16.Decode(buf[posStrBegin:idx]))

			result = append(result, s)
			posStrBegin = idx + 1
		}
	}

	return result
}

func GetAllVolumes() (volumes []string, err error) {
	buf := make([]uint16, windows.MAX_PATH) // this buffer should be well enough because volumeName is fixed length
	// cchBufferLength
	h, err := windows.FindFirstVolume(&buf[0], windows.MAX_PATH)
	if err != nil {
		return nil, err
	}
	defer windows.FindVolumeClose(h)

	for {
		v := syscall.UTF16ToString(buf)
		volumes = append(volumes, v)

		if err = windows.FindNextVolume(h, &buf[0], windows.MAX_PATH); err != nil {
			if err == windows.ERROR_NO_MORE_FILES {
				err = nil
			}
			break
		}
	}

	return
}

func GetAllDrives() (drives []string, err error) {

	// [in] nBufferLength
	// The maximum size of the buffer pointed to by lpBuffer, in TCHARs.
	// This size does not include the terminating null character.
	// If this parameter is zero, lpBuffer is not used.

	nBuf, _ := windows.GetLogicalDriveStrings(0, nil)

	buf := make([]uint16, nBuf+1)
	_, err = windows.GetLogicalDriveStrings(uint32(nBuf+1), &buf[0])

	if err != nil {
		return nil, err
	}

	drives = utf16ToStrings(buf)
	return
}
