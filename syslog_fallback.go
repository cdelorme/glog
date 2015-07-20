// +build !solaris,!linux,!darwin,!dragonfly,!freebsd,!netbsd,!openbsd

package log

func useSyslog() syslogger {
	return nil
}
