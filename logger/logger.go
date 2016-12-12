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

// NewLogger creates a new instance of a logger
func NewLogger() *Logger {
	// open a file
	f, err := os.OpenFile("logs/interview.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	// don't forget to close it
	defer f.Close()

	//log.SetOutput(f)

	log.Output(1, "this is an event")

	return &Logger{}
}

// LogMsg logs a message
func (l *Logger) LogMsg(msg string) {

}
