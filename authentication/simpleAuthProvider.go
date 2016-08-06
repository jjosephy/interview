package authentication

import (
	"errors"
)

// SimpleAuthProvider used for testing
type SimpleAuthProvider struct {
	SigningKey []byte
}

// AuthenticateUser authentication
func (p *SimpleAuthProvider) AuthenticateUser(name string, pwd string) (string, error) {

	if name == "tuser" && pwd == "fail" {
		return "", errors.New("invalid user")
	}

	return GenerateToken(p.SigningKey)
}

// ValidateToken validate
func (p *SimpleAuthProvider) ValidateToken(token string) (bool, error) {
	return ValidateToken(token, p.SigningKey)
}
