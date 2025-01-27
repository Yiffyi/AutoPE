//go:build generate

package native

//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output zsyscall_windows.go native.go
