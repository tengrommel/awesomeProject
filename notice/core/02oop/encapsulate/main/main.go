package main

import (
	"../model"
	"fmt"
)

func main() {
	p := model.NewPerson("smith")
	fmt.Println(p)
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(p)
	fmt.Println(p.Name, "age =", p.GetAge(), "sal =", p.GetSal())
}
