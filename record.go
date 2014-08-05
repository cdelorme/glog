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

type Record struct {
	Message    string
	Date       string
	File       string
	LineNumber int
	Level      LogLevel
	Args       []interface{}
}

func (record *Record) Build(message string, skip int, level LogLevel, args ...interface{}) {
	_, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		file = "???"
		line = 0
	}

	record.Message = message
	record.Date = time.Now().Format(time.Stamp)
	record.File = file
	record.LineNumber = line
	record.Level = level
	record.Args = args
}
