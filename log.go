package log

import (
	"fmt"
	"os"
	"sync"
)

const (
	ShortFormat = "[%s:%d]: %s\n" // [file:line]: message
	LongFormat  = "%s %10s %s %s" // timestamp severity application ShortFormat
)

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
)

const (
	Black = (iota + 30)
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

var Severities = []string{
	Debug:     "Debug",
	Info:      "Info",
	Notice:    "Notice",
	Warning:   "Warning",
	Error:     "Error",
	Critical:  "Critical",
	Alert:     "Alert",
	Emergency: "Emergency",
}

var Colors = []string{
	Debug:     fmt.Sprintf("\033[%dm", Cyan),
	Info:      fmt.Sprintf("\033[%dm", Green),
	Notice:    fmt.Sprintf("\033[%dm", Blue),
	Warning:   fmt.Sprintf("\033[%d;1m", Yellow),
	Error:     fmt.Sprintf("\033[%dm", Red),
	Critical:  fmt.Sprintf("\033[%d;1m", Red),
	Alert:     fmt.Sprintf("\033[%dm", Magenta),
	Emergency: fmt.Sprintf("\033[%d;1m", Magenta),
}

var mu sync.Mutex

func print(message string) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Fprintf(os.Stderr, message)
}
