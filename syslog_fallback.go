// +build windows,nacl,plan9

package log

func useSyslog() syslogger {
	return nil
}
