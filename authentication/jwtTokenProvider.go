package authentication

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

var (
    jwtDefaultKey []byte
    defaultKeyFunc    jwt.Keyfunc = func(t *jwt.Token) (interface{}, error) { return jwtDefaultKey, nil }
)

func GenerateToken(signingKey []byte) (string, error) {

    claims := jwt.StandardClaims{
        ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
        Issuer:    "nordstrom.net",
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, ex := token.SignedString(signingKey)

    if ex != nil {
        return "", ex
    }

    return tokenString, nil
}

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
