// +build !solaris,!linux,!darwin,!dragonfly,!freebsd

package log

func useColor() bool {
	return false
}
