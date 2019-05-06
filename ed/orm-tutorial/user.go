package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name string
	Email string
}

func InitialMigration()  {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()
	db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter, r *http.Request)  {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request)  {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]
	db.Create(&User{
		Name:name,
		Email:email,
	})
	fmt.Fprintf(w, "New User Successfully Created")
}

func DeleteUser(w http.ResponseWriter, r *http.Request)  {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()
	vars := mux.Vars(r)
	name := vars["name"]
	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)
	fmt.Fprintf(w, "New User Successfully Deleted")
}

func UpdateUser(w http.ResponseWriter, r *http.Request)  {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]
	var user User
	db.Where("name = ?", name).Find(&user)
	user.Email = email
	db.UpdateColumn(&user)
	fmt.Fprintf(w, "New User Successfully Updated")
}