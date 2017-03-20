// This package solves logging requirements by providing standardized
// severities and formatting, removing complex severity configuration,
// dealing with concurrency, properly using stderr, not closing your
// application on your behalf, and printing in color automatically
// when the terminal supports it.
//
// Examples
//
// To use the logger you may initialize a new instance:
//	logger := &glog.Logger{}
//	logger.Debug("this message requires debug log level")
//	logger.Error("this message will be printed with the default log level: %v\n", logger)
//
// Alternatively you would use the global instance:
//	glog.L.Debug("this message requires debug log level")
//	glog.L.Error("this message will be printed with the default log level: %v\n", logger)
//
// In both cases to print the debug messages simply set:
//	LOG_LEVEL=debug
package glog

import (
	"io"
	"os"
	"path"
	"strings"
	"sync"
)

// timestamp severity application [file:line]: message newline
const format = "%s %10s %s [%s:%d]: %s\n"

// severity (based on RFC-5424) of the message being written
// used when you need to
type Severity int

const (
	Debug Severity = iota
	Info
	Notice
	Warning
	Error
	Critical
	Alert
	Emergency
	silent
)

var severities = []string{
	Debug:     "debug",
	Info:      "info",
	Notice:    "notice",
	Warning:   "warning",
	Error:     "error",
	Critical:  "critical",
	Alert:     "alert",
	Emergency: "emergency",
	silent:    "silent",
}

var colors = []string{
	Debug:     "\033[36m",   // cyan
	Info:      "\033[32m",   // green
	Notice:    "\033[34m",   // blue
	Warning:   "\033[33;1m", // yellow
	Error:     "\033[31m",   // red
	Critical:  "\033[31;1m", // red bold
	Alert:     "\033[35m",   // magenta
	Emergency: "\033[35;1m", // magenta bold
	silent:    "",
}

var name string = path.Base(os.Args[0])
var stderr io.Writer = os.Stderr
var severity Severity = level()
var color bool = useColor()
var mu sync.Mutex

func level() Severity {
	l := os.Getenv("LOG_LEVEL")
	for s, v := range severities {
		if v == strings.ToLower(l) {
			return Severity(s)
		}
	}
	return Error
}

func print(msg string) {
	mu.Lock()
	stderr.Write([]byte(msg))
	mu.Unlock()
}
