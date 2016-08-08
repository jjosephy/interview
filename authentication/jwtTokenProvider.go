package authentication

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtDefaultKey  []byte
	defaultKeyFunc jwt.Keyfunc = func(t *jwt.Token) (interface{}, error) { return jwtDefaultKey, nil }
)

// GenerateToken creates a new JWT
func GenerateToken(signingKey []byte) (string, error) {

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 30).Unix(),
		Issuer:    "nordstrom.net",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, ex := token.SignedString(signingKey)

	if ex != nil {
		return "", ex
	}

	return tokenString, nil
}

// ValidateToken validates a provided JWT
func ValidateToken(token string, signingKey []byte) (bool, error) {
	var tk interface{}
	var err error

	tk, err = jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		time.Now().Add(time.Hour * 24).Unix()
		return signingKey, nil
	})

	if err != nil {
		return false, err
	}

	return tk.(*jwt.Token).Valid, nil
}
