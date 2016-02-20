package authentication

import (
    "errors"
)

type SimpleAuthProvider struct {
    SigningKey []byte
}

func (p *SimpleAuthProvider) AuthenticateUser(name string, pwd string) (string, error) {

    if name == "tuser" && pwd == "fail" {
        return "", errors.New("invalid user")
    }

    return GenerateToken(p.SigningKey)
}

func (p *SimpleAuthProvider) ValidateToken(token string) (bool, error) {
    return ValidateToken(token, p.SigningKey)
}
