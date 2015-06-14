package log

import (
	"fmt"
)

type Logger struct {
	Severity Severity
	Silent   bool
	NoColor  bool
	Syslog   bool
}

func (logger *Logger) log(record *Record) {
	if !logger.Silent && (record.Severity >= logger.Severity) {
		parsedMessage := record.Message
		if len(record.Args) > 0 {
			parsedMessage = fmt.Sprintf(parsedMessage, record.Args...)
		}
		if !logger.NoColor && useColor() {
			parsedMessage = fmt.Sprintf("%s%s\033[0m", Colors[record.Severity], parsedMessage)
		}
		parsedMessage = fmt.Sprintf(LogFormat, record.Date, Severities[record.Severity], record.File, record.LineNumber, parsedMessage)
		print(parsedMessage)
	}
}

func (logger *Logger) Emergency(message string, args ...interface{}) *Record {
	record := Record{}
	record.Build(message, 1, Emergency, args...)
	logger.log(&record)
	return &record
}

func (logger *Logger) Alert(message string, args ...interface{}) *Record {
	record := Record{}
	record.Build(message, 1, Alert, args...)
	logger.log(&record)
	return &record
}

func (logger *Logger) Critical(message string, args ...interface{}) *Record {
	record := Record{}
	record.Build(message, 1, Critical, args...)
	logger.log(&record)
	return &record
}

func (logger *Logger) Error(message string, args ...interface{}) *Record {
	record := Record{}
	record.Build(message, 1, Error, args...)
	logger.log(&record)
	return &record
}

func (logger *Logger) Warning(message string, args ...interface{}) *Record {
	record := Record{}
	record.Build(message, 1, Warning, args...)
	logger.log(&record)
	return &record
}

func (logger *Logger) Notice(message string, args ...interface{}) *Record {
	record := Record{}
	record.Build(message, 1, Notice, args...)
	logger.log(&record)
	return &record
}

func (logger *Logger) Info(message string, args ...interface{}) *Record {
	record := Record{}
	record.Build(message, 1, Info, args...)
	logger.log(&record)
	return &record
}

func (logger *Logger) Debug(message string, args ...interface{}) *Record {
	record := Record{}
	record.Build(message, 1, Debug, args...)
	logger.log(&record)
	return &record
}
