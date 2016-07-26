package log

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

type Logger struct{}

func (self *Logger) Log(s Severity, msg string, args ...interface{}) {
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

func (self *Logger) Emergency(message string, args ...interface{}) {
	self.Log(Emergency, message, args...)
}

func (self *Logger) Alert(message string, args ...interface{}) {
	self.Log(Alert, message, args...)
}

func (self *Logger) Critical(message string, args ...interface{}) {
	self.Log(Critical, message, args...)
}

func (self *Logger) Error(message string, args ...interface{}) {
	self.Log(Error, message, args...)
}

func (self *Logger) Warning(message string, args ...interface{}) {
	self.Log(Warning, message, args...)
}

func (self *Logger) Notice(message string, args ...interface{}) {
	self.Log(Notice, message, args...)
}

func (self *Logger) Info(message string, args ...interface{}) {
	self.Log(Info, message, args...)
}

func (self *Logger) Debug(message string, args ...interface{}) {
	self.Log(Debug, message, args...)
}
