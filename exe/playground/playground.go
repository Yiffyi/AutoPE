package main

import (
	"encoding/xml"
	"fmt"
	"runtime"
	"strings"
	"syscall"

	"github.com/yiffyi/autope/native"
	"golang.org/x/sys/windows"
)

func main() {
	runtime.LockOSThread()
	wimPath, err := syscall.UTF16PtrFromString("D:\\sources\\install.wim")
	fmt.Println(err)
	hImage, err := native.WIMCreateFile(wimPath, windows.GENERIC_READ, windows.OPEN_EXISTING, native.WIM_UNDOCUMENTED_BULLSHIT|native.WIM_FLAG_VERIFY, native.WIM_COMPRESS_NONE, nil)
	fmt.Println(hImage, err)

	wimInfo, err := native.WIMGetAttributes(hImage)
	fmt.Println("wimInfo:", wimInfo, err)

	xmlImageInfo, err := native.WIMGetImageInformation(hImage)
	fmt.Println("WIMGetImageInformation:", xmlImageInfo, err)

	var mapImageInfo map[string]interface{}
	err = xml.NewDecoder(strings.NewReader(xmlImageInfo)).Decode(&mapImageInfo)
	fmt.Println("WIMGetImageInformation:", mapImageInfo, err)
}
