package log

import (
	"fmt"
	"os"
	"sync"
)

type Logger struct {
	Level   LogLevel
	Silent  bool
	NoColor bool
	mu      sync.RWMutex
}

func (logger *Logger) log(record *Record) {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	if !logger.Silent && (record.Level >= logger.Level) {
		parsedMessage := record.Message
		if len(record.Args) > 0 {
			parsedMessage = fmt.Sprintf(parsedMessage, record.Args...)
		}
		if !logger.NoColor {
			parsedMessage = fmt.Sprintf("%s%s\033[0m", Colors[record.Level], parsedMessage)
		}
		parsedMessage = fmt.Sprintf(LogFormat, record.Date, LevelNames[record.Level], record.File, record.LineNumber, parsedMessage)
		fmt.Fprintln(os.Stderr, parsedMessage)
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
