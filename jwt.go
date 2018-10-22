package main

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(user string, target string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":   user,
		"target": target,
	})
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}
	return tokenString
}

func isValidToken(tokenString string) bool {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// claims
		return true
	} else {
		return false
	}
}
