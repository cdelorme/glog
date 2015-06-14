// +build darwin dragonfly freebsd netbsd openbsd

package log

import "syscall"

// override constant for alternate systems
const ioctlReadTermios = syscall.TIOCGETA
