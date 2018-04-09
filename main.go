package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jjosephy/interview/environment"
	"github.com/jjosephy/interview/handler"
	"github.com/jjosephy/interview/logger"
)

// Main entry point used to set up routes
func main() {
	logger.NewLogger()
	logger.LogInstance.LogMsg("Service Started")
	// TODO: try to find out how to get this called. Doesnt seem to be called on the defer
	defer logger.LogInstance.Close("Service End Signal Sent")

	// Must have config file present to run
	config, ex := ioutil.ReadFile("config/env.json")
	if ex != nil {
		panic(ex)
	}

	env := environment.NewEnvironment(config)
	p := env.AuthenticationProvider
	repo := env.Repository

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(env.WebPath)))
	mux.HandleFunc("/interview", handler.InterviewHandler(repo, p))
	mux.HandleFunc("/token", handler.TokenHandler(p))
	mux.HandleFunc("/user", handler.UserHandler(repo, p))
	err := http.ListenAndServeTLS(env.Port, env.PublicKey, env.PrivateKey, mux)
	if err != nil {
		fmt.Printf("main(): %s\n", err)
	}
}
