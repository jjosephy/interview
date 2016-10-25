package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jjosephy/interview/environment"
	"github.com/jjosephy/interview/handler"
)

// Main entry point used to set up routes
func main() {

	config, ex := ioutil.ReadFile("config/env.json")
	if ex != nil {
		panic(ex)
	}

	env := environment.NewEnvironment(config)

	fmt.Println("Starting Service")

	p := env.AuthenticationProvider
	repo := env.Repository

	mux := http.NewServeMux()
	// TODO: figure out path and a better way to configure
	mux.Handle("/", http.FileServer(http.Dir("../src/github.com/jjosephy/interview/web")))
	mux.HandleFunc("/interview", handler.InterviewHandler(repo, p))
	mux.HandleFunc("/token", handler.TokenHandler(p))
	err := http.ListenAndServeTLS(env.Port, env.PublicKey, env.PrivateKey, mux)
	if err != nil {
		fmt.Printf("main(): %s\n", err)
	}
}
