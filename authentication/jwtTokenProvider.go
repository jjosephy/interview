package authentication

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

var (
    jwtTestDefaultKey []byte
    defaultKeyFunc    jwt.Keyfunc = func(t *jwt.Token) (interface{}, error) { return jwtTestDefaultKey, nil }
)

func GenerateToken(signingKey []byte) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)
    token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
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
        return signingKey, nil
    })

    if err != nil {
        return false, err
    }

    return tk.(*jwt.Token).Valid, nil
}
