package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var mySigningKey = []byte("mysupersecreteparse")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodES256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "Elliot Forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Printf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, err
}

func main() {
	fmt.Println("My Simple Client")
	tokenString, err := GenerateJWT()
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println(tokenString)
}
