// +build !solaris,!linux,!darwin,!dragonfly,!freebsd

package glog

func useColor() bool {
	return false
}
