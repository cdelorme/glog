// +build darwin dragonfly freebsd linux,!appengine netbsd openbsd

package log

import (
	"os"
	"syscall"
	"unsafe"
)

// if stderr is a terminal, use color
func useColor() bool {
	var termios syscall.Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(os.Stderr.Fd()), ioctlReadTermios, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
}
