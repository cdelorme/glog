// +build !solaris,!linux,!darwin,!dragonfly,!freebsd,!windows

package log

func useColor() bool {
	return false
}
