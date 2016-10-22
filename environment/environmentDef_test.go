package environment

import (
    "testing"
    "github.com/jjosephy/interview/repository"
)

func Test_Success_CreateNewEnvironment(t *testing.T) {
    x := "Debug"
    p := "Path"
    r := &repository.MemoryInterviewRepository{}

    e := NewEnvironment(p, x, r)

    if e.Path != p {
        t.Fatal("Environment Path not set correctly")
    }

    if e.Type != x {
        t.Fatal("Environment Type not set correctly")
    }
}
