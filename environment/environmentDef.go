package environment

import (
	"encoding/json"
	"io/ioutil"

	"github.com/jjosephy/interview/authentication"
	"github.com/jjosephy/interview/repository"
)

type Environment struct {
	AuthenticationProvider authentication.Provider
	Port string
	PrivateKey string
	PublicKey string
	Repository repository.InterviewRepository
	Type string
}

func NewEnvironment(config []byte, ) (*Environment) {

	var i map[string]interface{}
	json.Unmarshal(config, &i)

	var signingKey []byte
	var e error
	if signingKey, e = ioutil.ReadFile(i["publicKey"].(string)); e != nil {
		panic(e)
	}

	return &Environment {
		AuthenticationProvider: authentication.NewAuthenticationProvder(i["authProvider"].(string), signingKey),
		Port: i["port"].(string),
		PrivateKey: i["privateKey"].(string),
		PublicKey: i["publicKey"].(string),
		Repository: repository.NewRepository(i["repository"].(string)),
		Type: i["type"].(string),
	}
}
