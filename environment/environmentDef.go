package environment

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/jjosephy/interview/authentication"
	"github.com/jjosephy/interview/repository"
)

// Environment is the main environment definition type
type Environment struct {
	AuthenticationProvider authentication.Provider
	LogPath                string
	Port                   string
	PrivateKey             string
	PublicKey              string
	Repository             repository.InterviewRepository
	Type                   string
	WebPath                string
}

// NewEnvironment creats a new instance of an environment
func NewEnvironment(config []byte) *Environment {

	var i map[string]interface{}
	// TODO: check for error or nil map
	json.Unmarshal(config, &i)

	var signingKey []byte
	var e error
	if signingKey, e = ioutil.ReadFile(i["publicKey"].(string)); e != nil {
		log.Panicf("error getting public key %s", e)
		panic(e)
	}

	return &Environment{
		AuthenticationProvider: authentication.NewAuthenticationProvder(i["authProvider"].(string), signingKey),
		LogPath:                i["logPath"].(string),
		Port:                   i["port"].(string),
		PrivateKey:             i["privateKey"].(string),
		PublicKey:              i["publicKey"].(string),
		Repository:             repository.NewRepository(i["repository"].(string)),
		Type:                   i["type"].(string),
		WebPath:                i["webpath"].(string),
	}
}
