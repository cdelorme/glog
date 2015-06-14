package log

import (
	"runtime"
	"time"
)

type Record struct {
	Message    string
	Date       string
	File       string
	LineNumber int
	Level      Severity
	Args       []interface{}
}

func (record *Record) Build(message string, skip int, level Severity, args ...interface{}) {
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
