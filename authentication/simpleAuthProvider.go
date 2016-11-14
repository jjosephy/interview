package authentication

import (
	"errors"
)

// SimpleAuthProvider used for testing
type SimpleAuthProvider struct {
	SigningKey []byte
	users      map[string]string
}

// NewSimpleAuthProvider returns a new instance of SimpleAuthProvider
func NewSimpleAuthProvider(key []byte) *SimpleAuthProvider {
	m := make(map[string]string)
	m["user"] = "password"

	return &SimpleAuthProvider{
		SigningKey: key,
		users:      m,
	}
}

// AuthenticateUser authentication
func (p *SimpleAuthProvider) AuthenticateUser(name string, pwd string) (string, error) {
	if val, ok := p.users[name]; ok {
		if val == pwd {
			return GenerateToken(p.SigningKey)
		}
	}

	return "", errors.New("invalid user")
}

// ValidateToken validate
func (p *SimpleAuthProvider) ValidateToken(token string) (bool, error) {
	return ValidateToken(token, p.SigningKey)
}
