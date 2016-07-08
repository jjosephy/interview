package authentication

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

var (
    jwtDefaultKey []byte
    defaultKeyFunc    jwt.Keyfunc = func(t *jwt.Token) (interface{}, error) { return jwtDefaultKey, nil }
)

type TokenTimeClaims struct {
    exp int64
    jwt.StandardClaims
}

func GenerateToken(signingKey []byte) (string, error) {

    claims := TokenTimeClaims {
        time.Now().Add(time.Hour * 24).Unix(),
        jwt.StandardClaims{
            ExpiresAt: 15000,
            Issuer:    "nordstrom.net",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    //token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
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
