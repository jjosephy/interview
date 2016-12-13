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

// NewLogger creates a new instance of a logger
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
	t := time.Now()
	log.Printf("%s | INFO | %s", t.UTC().Format(time.UnixDate), msg)
}

// LogError logs a message
func (l *Logger) LogError(msg string) {
	t := time.Now()
	log.Printf("%s | ERROR | %s", t.UTC().Format(time.UnixDate), msg)
}

// Close closes the Logger
func (l *Logger) Close() {
	e := l.f.Close()
	if e != nil {
		fmt.Printf("Error trying to close file %s\n", e)
	}
}
