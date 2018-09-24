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
	fmt.Println("Go MySQL Tutorial")
	db, err := sql.Open("mysql", "zhimaa:zhimaa@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	results, err := db.Query("SELECT name from users")
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
