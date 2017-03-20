package glog

// declare constant manually since syscall may be unreliable
const ioctlReadTermios = 0x5401 // syscall.TCGETS
