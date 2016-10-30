package environment

import (
    "testing"
)

const j = `{
			"authProvider": "SimpleAuthProvider",
			"port": ":8443",
			"publicKey" : "./cert.pem",
			"privateKey" : "./private_key",
			"repository" : "MemoryRepository",
			"type" : "debug",
			"webpath" : "/Users/jjosephy/Source/go/src/github.com/jjosephy/interview/web",
			}`

func Test_Success_CreateNewEnvironment(t *testing.T) {
	//e := NewEnvironment([]byte(j))

	/*
	if e.Port != ":8443" {
		t.Fatal("Invalid Port")
	}

	t.Logf("%v", e)
	*/
}
