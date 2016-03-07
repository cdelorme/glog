package log

import (
	"io"
	"os"
	"path"
	"strings"
	"sync"
)

const format = "%s %10s %s [%s:%d]: %s\n" // timestamp severity application [file:line]: message newline

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
	Debug:     "Debug",
	Info:      "Info",
	Notice:    "Notice",
	Warning:   "Warning",
	Error:     "Error",
	Critical:  "Critical",
	Alert:     "Alert",
	Emergency: "Emergency",
	silent:    "Silent",
}

var colors = []string{
	Debug:     "\033[36m",   // Cyan
	Info:      "\033[32m",   // Green
	Notice:    "\033[34m",   // Bluew
	Warning:   "\033[33;1m", // Yellow
	Error:     "\033[31m",   // Red
	Critical:  "\033[31;1m", // Red Bold
	Alert:     "\033[35m",   // Magenta
	Emergency: "\033[35;1m", // Magenta Bold
	silent:    "",
}

var name string
var color bool
var severity Severity
var mu sync.Mutex
var stderr io.Writer = os.Stderr

func init() {
	name = path.Base(os.Args[0])
	color = useColor()
	severity = level()
}

func level() Severity {
	l := os.Getenv("LOG_LEVEL")
	for s, v := range severities {
		if strings.ToLower(v) == strings.ToLower(l) {
			return Severity(s)
		}
	}
	return Error
}

func print(msg string) {
	mu.Lock()
	defer mu.Unlock()
	stderr.Write([]byte(msg))
}
