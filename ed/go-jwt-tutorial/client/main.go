package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//var mySigningKey = os.Getenv("MY_JWT_TOKEN")
var mySigningKey = []byte("mysupersecretphrase")

func homePage(w http.ResponseWriter, r *http.Request)  {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:9009", nil)
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, string(body))
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "Elliot Forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Printf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func handleRequests()  {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Println("My Simple Client")
	handleRequests()
}
