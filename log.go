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

func GetLogger(logTo string, logFile string, mongoDatabase string, mongoUsername string, mongoPassword string) Logger {
	var logger Logger

	if logTo == "file" && logFile != "" {
		logger = &FileLogger{File: logFile}
	} else if logTo == "mongo" && mongoDatabase != "" {
		logger = &MongoLogger{Database: mongoDatabase, Username: mongoUsername, Password: mongoPassword}
	} else {
		logger = &StandardLogger{}
	}

	return logger
}
