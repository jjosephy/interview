package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Logger main type for logging
type Logger struct {
	f *os.File
}

// LogInstance singleton instance in package
var LogInstance *Logger

// NewLogger creates a new instance of a logger with a standard log schema
// [0] - Time Stamp
// [1] - INFO | ERROR | WARNING
// [2] - Custom Message
// [3] - Http Status Code
func NewLogger() {
	// open a file
	f, err := os.OpenFile("logs/interview.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		defer f.Close()
	}
	log.SetOutput(f)

	LogInstance = &Logger{f: f}
}

// LogMsg logs a message
func (l *Logger) LogMsg(msg string) {
	// Add HTTPStatus code
	l.format("INFO", msg)
}

// LogError logs a message
func (l *Logger) LogError(msg string) {
	// Add Service Error Code
	l.format("ERROR", msg)
}

// Close closes the Logger
func (l *Logger) Close(msg string) {
	if len(msg) > 0 {
		l.LogMsg(msg)
	}

	e := l.f.Close()
	if e != nil {
		fmt.Printf("Error trying to close file %s\n", e)
	}
}

func (l *Logger) format(msgType string, msg string) {
	t := time.Now()
	log.Printf("%s | %s | %s", t.UTC().Format(time.UnixDate), msgType, msg)
}
