package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/couchbase/gocb.v1"
)

type Account struct {
	Type     string `json:"type,omitempty"`
	UserName string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func main() {
	cluster, _ := gocb.Connect("couchbase://localhost")
	bucket, _ := cluster.OpenBucket("example", "")
	hash, _ := bcrypt.GenerateFromPassword([]byte("my-password"), 10)
	account := Account{
		Type:     "account",
		UserName: "nraboy",
		Password: string(hash),
	}
	bucket.Insert(account.UserName, account, 0)
	account = Account{}
	bucket.Get("nraboy", &account)
	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte("my-password"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
