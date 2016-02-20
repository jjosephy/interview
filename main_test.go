package main

import (
    "errors"
    "testing"
)


type DelegatingHandler func(i int) (bool, int, error)

func handleInternal(i int)(bool, int, error) {
    if i == 20 {
        return false, i, errors.New("match cannot be 20")
    }

    return true, (i * 2), nil
}

func th(t *testing.T, dh DelegatingHandler) {
    b, x, err := dh(10)

    if err != nil {
        t.Fatal("failed")
    }

    if b != true {
        t.Fatal("Failed")
    }

    b2, _, err2 := dh(x)

    if err2 == nil {
        t.Fatal("failed")
    } else {
        t.Log("err: ", err2)
    }

    if b2 != false {
        t.Fatal("Failed")
    }
}

func Test_DelegatingHandlers(t *testing.T) {
    // go test -v interview -run Test_DelegatingHandlers
    var f DelegatingHandler
    f = handleInternal
    th(t, f)
}
