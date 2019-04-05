package main

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	ID        string `validate:"omitempty,uuid"`
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	UserName  string `validate:"required,email"`
	Password  string `validate:"required,gte=10"`
	Type      string `validate:"isdefault"`
}

func main() {
	user := User{
		UserName: "test@test.com",
		Password: "1234567890",
	}
	validate := validator.New()
	err := validate.StructExcept(user, "FirstName", "LastName")
	if err != nil {
		fmt.Println(err.Error())
	}
	password := "12345"
	err = validate.Var(password, "required,gte=10")
	if err != nil {
		fmt.Println(err.Error())
	}
}
