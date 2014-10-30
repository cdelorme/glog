package log

import (
    "fmt"
)

type Logger struct {
    Level   LogLevel
    Silent  bool
    NoColor bool
}

func (logger *Logger) log(record *Record, color Color) {

    if !logger.Silent && (logger.Level == DEFAULT || record.Level >= logger.Level) {
    	parsedMessage := record.Message

    	if len(record.Args) > 0 {
    		parsedMessage = fmt.Sprintf(record.Message, &record.Args)
    	}

    	parsedMessage = fmt.Sprintf("[%s] %s:%d, %s", record.Date, record.File, record.LineNumber, parsedMessage)

    	if !logger.NoColor {
    		parsedMessage = fmt.Sprintf("%s[%dm%s%s[%dm", ESCAPE, color, parsedMessage, ESCAPE, RESET)
    	}

    	fmt.Println(parsedMessage)
    }
}

func (logger *Logger) Emergency(message string, args ...interface{}) *Record {
    record := Record{}
    record.Build(message, 1, EMERGENCY, args...)

    logger.log(&record, RED)

    return &record
}

func (logger *Logger) Alert(message string, args ...interface{}) *Record {
    record := Record{}
    record.Build(message, 1, ALERT, args...)

    logger.log(&record, RED)

    return &record
}

func (logger *Logger) Critical(message string, args ...interface{}) *Record {
    record := Record{}
    record.Build(message, 1, CRITICAL, args...)

    logger.log(&record, MAGENTA)

    return &record
}

func (logger *Logger) Error(message string, args ...interface{}) *Record {
    record := Record{}
    record.Build(message, 1, ERROR, args...)

    logger.log(&record, MAGENTA)

    return &record
}

func (logger *Logger) Notice(message string, args ...interface{}) *Record {
    record := Record{}
    record.Build(message, 1, NOTICE, args...)

    logger.log(&record, YELLOW)

    return &record
}

func (logger *Logger) Warning(message string, args ...interface{}) *Record {
    record := Record{}
    record.Build(message, 1, WARNING, args...)

    logger.log(&record, YELLOW)

    return &record
}

func (logger *Logger) Info(message string, args ...interface{}) *Record {
    record := Record{}
    record.Build(message, 1, INFO, args...)

    logger.log(&record, RESET)

    return &record
}

func (logger *Logger) Debug(message string, args ...interface{}) *Record {
    record := Record{}
    record.Build(message, 1, DEBUG, args...)

    logger.log(&record, BLUE)

    return &record
}
