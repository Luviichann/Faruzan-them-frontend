package tools

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	Key string
	jwt.RegisteredClaims
}

func ReleaseToken(key string) (string, error) {
	expirationTime := time.Now().Add(7 * 86400 * time.Second)
	claims := Claims{
		Key: key,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "faruzan.cn",
			Subject:   "token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
