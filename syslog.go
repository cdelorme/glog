// +build !windows,!nacl,!plan9

package log

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const SyslogFormat = "<%d>%s %s[%d]: %s" // <PRI>TIMESTAMP HOSTNAME TAG[PID]: MSG
const severityMask = 0x07
const facilityMask = 0xf8

type syslog struct {
	tag  string
	host string

	conn net.Conn
	mu   sync.Mutex
}

func (self *syslog) init() {
	self.mu.Lock()
	defer self.mu.Unlock()
	if tag, err := filepath.Abs(os.Args[0]); err != nil {
		self.tag = os.Args[0]
	} else {
		self.tag = tag
	}
	hostname, _ := os.Hostname()
	if hostname == "" {
		hostname = "localhost"
	}
	self.host = hostname
	self.connect()
}

func (self *syslog) Print(severity Severity, message string) {
	// @todo(casey): verify conversion to integer value for calculation below
	s := ((7 - int(severity)) & facilityMask) | ((7 - int(severity)) & severityMask)

	self.mu.Lock()
	defer self.mu.Unlock()

	if err := self.write(s, message); err == nil {
		return
	} else {
		self.connect()
	}
	self.write(s, message)
}

func (self *syslog) write(severity int, message string) error {
	_, err := fmt.Fprintf(self.conn,
		SyslogFormat,
		severity,
		time.Now().Format(time.Stamp),
		self.tag,
		self.host,
		os.Getpid(),
		message)
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
