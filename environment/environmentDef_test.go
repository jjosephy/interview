package environment

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

const j = `{
    "authProvider": "SimpleAuthProvider",
    "logPath": "/home/jjosephy/Source/go/bin/logs",
    "port": ":8443",
    "publicKey" : "./cert.pem",
    "privateKey" : "./private_key",
    "repository" : "MemoryRepository",
    "type" : "debug",
    "webpath" : "/home/jjosephy/Source/go/src/github.com/jjosephy/interview/web"
}`

func init() {

	var (
		cmdOut []byte
		err    error
	)
	cmdName := "openssl"
	cmdArgs := []string{"genrsa", "-out", "private_key", "2048"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running openssl to create private key: ", err, cmdOut)
		os.Exit(1)
	}

	cmd := fmt.Sprintln("openssl req -new -x509 -key private_key -out cert.pem",
		"-days 365 -subj '/C=US/ST=Washington/L=Seattle/O=Comp/CN=localhost'")
	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Printf("error trying to cmd %s %s", out, err)
		os.Exit(1)
	}
}

func clean() {
	cmd := fmt.Sprintln("rm -f cert.pem private_key")
	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Printf("error trying to clean up certs %s %s", out, err)
	}
}

func Test_Success_CreateNewEnvironment(t *testing.T) {
	e := NewEnvironment([]byte(j))

	if e.Port != ":8443" {
		t.Fatal("Invalid Port")
	}

	t.Logf("%v", e)

	defer clean()
}
