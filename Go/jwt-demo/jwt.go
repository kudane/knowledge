package main

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("secret key na jaaa")

type CustomClaims struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func SignToken(claims *CustomClaims) string {
	genToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := genToken.SignedString(secret)
	if err != nil {
		panic("error na ja")
	}
	return tokenString
}

func IsTokenValid(tokenString string) bool {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	return token.Valid
}

func IsTokenExpired(tokenString string) bool {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	return errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet)
}
