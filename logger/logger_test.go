package logger

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

const (
	TestToken1 = "{ID1}"
)

func init() {
	//make logs directory. check if it already exists.
	if e := os.Mkdir("logs", os.ModeDir); e != nil {
		if os.IsExist(e) {
			fmt.Printf("logs directory already exists\n")
		} else {
			fmt.Printf("Error trying to create log directory. : %s\n", e)
		}
	}
}

func cleanup() {
	// delete the log directory
	if e := os.RemoveAll("logs"); e != nil {
		fmt.Printf("error deleting logs directory %s \n", e)
	}
}

func Test_BasicLogging(t *testing.T) {
	NewLogger()
	LogInstance.LogMsg(fmt.Sprintf("Test Messsage %s", TestToken1))
	defer cleanup()
	defer LogInstance.Close("--End Test--")

	f, err := os.Open("logs/interview.log")
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	r := bufio.NewReader(f)
	b := false
	for err != io.EOF {
		line, err := r.ReadString(10)
		if err != nil {
			break
		}
		sp := strings.Split(line, "|")
		if strings.Contains(sp[2], TestToken1) {
			b = true
			break
		}
	}

	f.Close()

	if !b {
		t.Fatalf("Could not parse test message\n")
	}
}
