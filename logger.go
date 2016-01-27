package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var app string

func init() {
	app = os.Args[0]
	if t, err := filepath.Abs(os.Args[0]); err == nil {
		app = t
	}
}

type syslogger interface {
	Print(severity Severity, message string)
}

type Logger struct {
	Severity Severity
	Silent   bool
	color    bool
	syslog   syslogger
}

func (self *Logger) log(message string, severity Severity, args ...interface{}) {
	if self.Silent || severity < self.Severity {
		return
	}

	m := message
	if len(args) > 0 {
		m = fmt.Sprintf(m, args...)
	}
	mu.Lock()
	if self.color && self.syslog == nil {
		m = fmt.Sprintf("%s%s\033[0m", Colors[severity], m)
	}
	mu.Unlock()

	_, file, line, _ := runtime.Caller(2)
	m = fmt.Sprintf(ShortFormat, filepath.Base(file), line, m)

	if self.syslog == nil {
		print(fmt.Sprintf(LongFormat, time.Now().Format(time.Stamp), app, Severities[severity], m))
	} else {
		self.syslog.Print(severity, m)
	}
}

func (self *Logger) Log(message string, severity Severity, args ...interface{}) {
	self.log(message, severity, args...)
}

func (self *Logger) Emergency(message string, args ...interface{}) {
	self.log(message, Emergency, args...)
}

func (self *Logger) Alert(message string, args ...interface{}) {
	self.log(message, Alert, args...)
}

func (self *Logger) Critical(message string, args ...interface{}) {
	self.log(message, Critical, args...)
}

func (self *Logger) Error(message string, args ...interface{}) {
	self.log(message, Error, args...)
}

func (self *Logger) Warning(message string, args ...interface{}) {
	self.log(message, Warning, args...)
}

func (self *Logger) Notice(message string, args ...interface{}) {
	self.log(message, Notice, args...)
}

func (self *Logger) Info(message string, args ...interface{}) {
	self.log(message, Info, args...)
}

func (self *Logger) Debug(message string, args ...interface{}) {
	self.log(message, Debug, args...)
}

func (self *Logger) Color() bool {
	mu.Lock()
	defer mu.Unlock()
	self.color = useColor()
	return self.color
}

func (self *Logger) Syslog() bool {
	self.syslog = useSyslog()
	mu.Lock()
	defer mu.Unlock()
	if self.syslog != nil {
		self.color = false
	}
	return self.syslog != nil
}
