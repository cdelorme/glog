// +build darwin dragonfly freebsd netbsd openbsd

package glog

import "syscall"

// override constant for alternate systems
const ioctlReadTermios = syscall.TIOCGETA
