package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	fmt.Println("Go Mysql")
	db, err := sql.Open("mysql", "zhimaa:zhimaa@tcp(127.0.0.1:3306)/fantest")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Successfully Connected to Mysql database")
	insert, err := db.Query("INSERT INTO users VALUE ('teng')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Successfully insert to user table")
	results, err := db.Query("SELECT name FROM  users")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var user User
		err = results.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.Name)
	}
}
