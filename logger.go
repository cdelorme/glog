package log

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

type Logger struct{}

func (self *Logger) log(msg string, s Severity, args ...interface{}) {
	if s < severity {
		return
	}

	m := msg
	if len(args) > 0 {
		m = fmt.Sprintf(m, args...)
	}

	_, f, l, _ := runtime.Caller(2)

	if color {
		m = fmt.Sprintf(colors[s]+format, time.Now().Format(time.Stamp), severities[s], name, path.Base(f), l, "\033[0m"+m)
	} else {
		m = fmt.Sprintf(format, time.Now().Format(time.Stamp), severities[s], name, path.Base(f), l, m)
	}

	print(m)
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
