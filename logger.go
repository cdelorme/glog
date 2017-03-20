package glog

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

// package-global Logger instance for logging directly
var L Logger

// a logger that combines fmt parameters with RFC-5424 compatible log levels
// and writes to stderr in a concurrently safe way, while pushing severity
// controls into the hands of the user so the application remains simple
type Logger struct{}

// a direct call to log a message, where you supply the severity directly
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

// emergency is not adviced for use by applications
func (self *Logger) Emergency(msg string, args ...interface{}) {
	self.Log(Emergency, msg, args...)
}

// alert should probably be used to invoke notifications (eg. via email or SMS)
func (self *Logger) Alert(msg string, args ...interface{}) {
	self.Log(Alert, msg, args...)
}

// critical generally means something outside the application caused a failure
// and is best matched with panic catching events
func (self *Logger) Critical(msg string, args ...interface{}) {
	self.Log(Critical, msg, args...)
}

// errors are generally known quantities within the application that you have
// planned for, but may wish to record the occurrence
func (self *Logger) Error(msg string, args ...interface{}) {
	self.Log(Error, msg, args...)
}

// warning is generally used when something non-essential fails or when
// your application anticipates future problems ahead
func (self *Logger) Warning(msg string, args ...interface{}) {
	self.Log(Warning, msg, args...)
}

// notice may be used to share unusual events that do not disrupt execution
func (self *Logger) Notice(msg string, args ...interface{}) {
	self.Log(Notice, msg, args...)
}

// info is generally helpful output regarding and to identify normal operation
func (self *Logger) Info(msg string, args ...interface{}) {
	self.Log(Info, msg, args...)
}

// debug is development output that is only meant for debugging the application
func (self *Logger) Debug(msg string, args ...interface{}) {
	self.Log(Debug, msg, args...)
}
