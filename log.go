package log

import (
	"fmt"
	"os"
	"sync"
)

type Severity int

const LogFormat = "%s %10s [%s:%d] %s"

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

var Severities = []string{
	Debug:     "Debug",
	Info:      "Info",
	Notice:    "Warning",
	Warning:   "Notice",
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
	fmt.Fprintln(os.Stderr, message)
}
