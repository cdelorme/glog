package log

type LogLevel int
type Color int

const ESCAPE = "\x1b"
const RESET = 0

const (
	BLACK Color = (iota + 30)
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
)

const SYSLOG = "%s %s %s:%d, %s"

const (
	DEFAULT LogLevel = iota
	EMERGENCY
	ALERT
	CRITICAL
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
)

var LevelNames = []string{
	"Debug",
	"Emergency",
	"Alert",
	"Critical",
	"Error",
	"Notice",
	"Warning",
	"Info",
	"Debug",
}
