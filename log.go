package log

/**
 * Basic logging /w color support
 *
 * Author: Casey Delow <cdelorme@gmail.com>
 * Date: 2014-8-5
 */

type LogLevel int
type Color int

const ESCAPE = "\x1b"
const RESET = 0

const (
	BLACK Color = iota + 30
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
)

const (
	EMERGENCY LogLevel = iota
	ALERT
	CRITICAL
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
)

var LevelNames = []string{
	"Emergency",
	"Alert",
	"Critical",
	"Error",
	"Notice",
	"Warning",
	"Info",
	"Debug",
}
