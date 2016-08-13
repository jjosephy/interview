package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jjosephy/interview/authentication"
	"github.com/jjosephy/interview/handler"
	"github.com/jjosephy/interview/repository"
)

const (
	// Port to use for the Service
	Port = ":8443"

	// PrivateKey is key to use for TLS
	PrivateKey = "./private_key"

	// PublicKey is key to use for TLS
	PublicKey = "./cert.pem"
)

// Main entry point used to set up routes
func main() {
	fmt.Println("Starting Service")
	var e error
	var signingKey []byte
	if signingKey, e = ioutil.ReadFile(PublicKey); e != nil {
		panic(e)
	}

	// TODO: figure out injection pattern and config
	p := authentication.SimpleAuthProvider{SigningKey: signingKey}
	repo := repository.New("mongodb://localhost")

	mux := http.NewServeMux()
	// TODO: figure out path and a better way to configure
	mux.Handle("/", http.FileServer(http.Dir("../src/github.com/jjosephy/interview/web")))
	mux.HandleFunc("/interview", handler.InterviewHandler(repo, &p))
	mux.HandleFunc("/token", handler.TokenHandler(&p))
	err := http.ListenAndServeTLS(Port, PublicKey, PrivateKey, mux)
	if err != nil {
		fmt.Printf("main(): %s\n", err)
	}
}
