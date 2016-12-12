package logger

import (
	"fmt"
	"log"
	"os"
)

// Logger main type for logging
type Logger struct {
	f *os.File
}

// LogInstance singleton instance in package
var LogInstance *Logger

// NewLogger creates a new instance of a logger
func NewLogger() {
	// open a file
	f, err := os.OpenFile("logs/interview.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	log.SetOutput(f)

	LogInstance = &Logger{f: f}
}

// LogMsg logs a message
func (l *Logger) LogMsg(msg string) {
	log.Output(1, msg)
}

// Close closes the Logger
func (l *Logger) Close() {
	l.f.Close()
}
