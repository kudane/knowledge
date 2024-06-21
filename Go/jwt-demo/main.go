package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	fmt.Println("Hello, World!")

	claims := CustomClaims{
		Id:   "1",
		Name: "ssss",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Second)),
		},
	}

	token := SignToken(&claims)

	fmt.Println(token)
	fmt.Println(IsTokenValid(token))
	fmt.Println(IsTokenExpired(token))
}
