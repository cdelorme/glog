package log

/**
 * Logger struct
 *  prints to console and returns LogMessage
 *
 * Author: Casey Delow <cdelorme@gmail.com>
 * Date: 2014-8-5
 */

import (
	"fmt"
)

type Logger struct {
	Level   LogLevel
	Silent  bool
	NoColor bool
}

func (logger *Logger) log(logMessage *LogMessage, color Color) {

	if !logger.Silent && logMessage.Level >= logger.Level {
		parsedMessage := logMessage.Message

		if len(logMessage.Args) > 0 {
			parsedMessage = fmt.Sprintf(logMessage.Message, &logMessage.Args)
		}

		parsedMessage = fmt.Sprintf("[%s] %s:%d, %s", logMessage.Date, logMessage.File, logMessage.LineNumber, parsedMessage)

		if !logger.NoColor {
			parsedMessage = fmt.Sprintf("%s[%dm%s%s[%dm", ESCAPE, color, parsedMessage, ESCAPE, RESET)
		}

		fmt.Println(parsedMessage)
	}
}

func (logger *Logger) Emergency(message string, args ...interface{}) *LogMessage {
	logMessage := LogMessage{}
	logMessage.Build(message, 1, EMERGENCY, args...)

	logger.log(&logMessage, RED)

	return &logMessage
}

func (logger *Logger) Alert(message string, args ...interface{}) *LogMessage {
	logMessage := LogMessage{}
	logMessage.Build(message, 1, ALERT, args...)

	logger.log(&logMessage, RED)

	return &logMessage
}

func (logger *Logger) Critical(message string, args ...interface{}) *LogMessage {
	logMessage := LogMessage{}
	logMessage.Build(message, 1, CRITICAL, args...)

	logger.log(&logMessage, MAGENTA)

	return &logMessage
}

func (logger *Logger) Error(message string, args ...interface{}) *LogMessage {
	logMessage := LogMessage{}
	logMessage.Build(message, 1, ERROR, args...)

	logger.log(&logMessage, MAGENTA)

	return &logMessage
}

func (logger *Logger) Notice(message string, args ...interface{}) *LogMessage {
	logMessage := LogMessage{}
	logMessage.Build(message, 1, NOTICE, args...)

	logger.log(&logMessage, YELLOW)

	return &logMessage
}

func (logger *Logger) Warning(message string, args ...interface{}) *LogMessage {
	logMessage := LogMessage{}
	logMessage.Build(message, 1, WARNING, args...)

	logger.log(&logMessage, YELLOW)

	return &logMessage
}

func (logger *Logger) Info(message string, args ...interface{}) *LogMessage {
	logMessage := LogMessage{}
	logMessage.Build(message, 1, INFO, args...)

	logger.log(&logMessage, RESET)

	return &logMessage
}

func (logger *Logger) Debug(message string, args ...interface{}) *LogMessage {
	logMessage := LogMessage{}
	logMessage.Build(message, 1, DEBUG, args...)

	logger.log(&logMessage, BLUE)

	return &logMessage
}
