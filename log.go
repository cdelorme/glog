package log

type LogLevel int
type Color int

const COLORESCAPE = "\033"
const COLORRESET = COLORESCAPE + "[0m"

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
