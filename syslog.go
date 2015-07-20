// +build !windows,!nacl,!plan9

package log

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"
)

const SyslogFormat = "<%d>%s %s[%d]: %s" // <PRI>TIMESTAMP TAG[PID]: MSG
const severityMask = 0x07
const facilityMask = 0xf8

type syslog struct {
	conn net.Conn
	tag  string
}

func (self *syslog) init() {
	mu.Lock()
	defer mu.Unlock()
	if tag, err := filepath.Abs(os.Args[0]); err != nil {
		self.tag = os.Args[0]
	} else {
		self.tag = tag
	}
	self.connect()
}

func (self *syslog) Print(severity Severity, message string) {
	s := ((7 - int(severity)) & facilityMask) | ((7 - int(severity)) & severityMask)

	mu.Lock()
	defer mu.Unlock()

	if err := self.write(s, message); err == nil {
		return
	} else {
		self.connect()
	}
	self.write(s, message)
}

func (self *syslog) write(severity int, message string) error {
	_, err := fmt.Fprintf(self.conn,
		SyslogFormat,                  // <PRI>TIMESTAMP TAG[PID]: MSG
		severity,                      // PRI
		time.Now().Format(time.Stamp), // TIMESTAMP
		self.tag,                      // TAG
		os.Getpid(),                   // PID
		message)                       // MSG
	return err
}

func (self *syslog) connect() {
	if self.conn != nil {
		self.conn = nil
	}

	logTypes := []string{"unixgram", "unix"}
	logPaths := []string{"/dev/log", "/var/run/syslog", "/var/run/log"}
	for _, network := range logTypes {
		for _, path := range logPaths {
			conn, err := net.Dial(network, path)
			if err != nil {
				continue
			}
			self.conn = conn
			return
		}
	}
}

func useSyslog() syslogger {
	s := &syslog{}
	s.init()
	return s
}
