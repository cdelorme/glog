// +build !windows,!plan9,!netbsd,!openbsd

package log

import (
	"os"
	"syscall"
	"unsafe"
)

func useColor() bool {
	var termios syscall.Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(os.Stderr.Fd()), ioctlReadTermios, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
}
