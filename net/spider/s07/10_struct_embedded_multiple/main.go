package main

import "fmt"

type info struct{ fName, lName string }

type person struct {
	info
	age, height int
}

type student struct {
	person
	id int
}

func main() {
	var s student
	s = student{person{info{"Jim", "Brown"}, 20, 178}, 1200}
	fmt.Printf("%v \n\n", s)
	fmt.Printf("%#v \n\n", s)
	s = student{
		person: person{
			info: info{fName: "Tim", lName: "Green"},
			age:  22,
		},
		id: 1400,
	}
	fmt.Printf("%v \n\n", s)
	fmt.Printf("%v \n", s)
}
