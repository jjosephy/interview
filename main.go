package main

import (
    "fmt"
    "github.com/jjosephy/go/interview/authentication"
    "github.com/jjosephy/go/interview/handler"
    "github.com/jjosephy/go/interview/repository"
    "io/ioutil"
    "net/http"
)

const (
    PORT       = ":8443"
    PRIV_KEY   = "./private_key"
    PUBLIC_KEY = "./cert.pem"
)

// Main entry point used to set up routes //
func main() {
    var e error
    var signingKey []byte
    if signingKey, e = ioutil.ReadFile(PUBLIC_KEY); e != nil {
        panic(e)
    }

    // TODO: figure out injection pattern and config
    p := authentication.SimpleAuthProvider { SigningKey : signingKey }
    repo := repository.DBInterviewRepository{ Uri: "mongodb://localhost" }

    mux := http.NewServeMux()
    // TODO: figure out path and a better way to configure
    mux.Handle("/", http.FileServer(http.Dir("../src/github.com/jjosephy/go/interview/web")))
    mux.HandleFunc("/interview", handler.InterviewHandler(&repo))
    mux.HandleFunc("/token", handler.TokenHandler(&p))
    err := http.ListenAndServeTLS(PORT, PUBLIC_KEY, PRIV_KEY, mux)
    if err != nil {
       fmt.Printf("main(): %s\n", err)
    }
}
