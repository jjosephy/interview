package logger

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

const (
	TestToken1 = "{ID1}"
)

func init() {
	cmd := fmt.Sprintln("[ ! -d logs/ ] && mkdir logs/")
	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		if len(out) > 0 {
			fmt.Printf("error trying to create the log %s %s", out, err)
			os.Exit(1)
		}
	}
}

func cleanup() {
	cmd := fmt.Sprintln("rm -f -r logs")
	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Printf("error trying to delete log directory %s %s", out, err)
	}
}

func Test_BasicLogging(t *testing.T) {
	NewLogger()
	LogInstance.LogMsg(fmt.Sprintf("Test Messsage %s", TestToken1))

	/*
		f, err := os.Open("logs/interview.log")
		if err != nil {
			fmt.Printf("error opening file: %v", err)
		}

		defer f.Close()
		r := bufio.NewReader(f)
		line, err := r.ReadString(10)
		for err != io.EOF {
			fmt.Print(line)
			line, err = r.ReadString(10)
		}
	*/

	defer cleanup()
	defer LogInstance.Close()
}
