package log

/**
 * LogMessage container for storage and transmission
 *
 * Author: Casey Delow <cdelorme@gmail.com>
 * Date: 2014-7-20
 */

import (
	"runtime"
	"time"
)

type LogMessage struct {
	Message    string
	Date       string
	File       string
	LineNumber int
	Level      LogLevel
	Args       []interface{}
}

func (logMessage *LogMessage) Build(message string, skip int, level LogLevel, args ...interface{}) {
	_, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		file = "???"
		line = 0
	}

	logMessage.Message = message
	logMessage.Date = time.Now().Format(time.Stamp)
	logMessage.File = file
	logMessage.LineNumber = line
	logMessage.Level = level
	logMessage.Args = args
}
