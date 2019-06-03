package main

import (
	"fmt"
	"github.com/tengrommel/awesomeProject/di/ch1/code-smells/code-bloat"
)

type a struct {
	Id int64
	Name string
}

func (a a)AddId() a {
	a.Id = a.Id + 1
	return a
}

type aList []a

func (a aList)dd()  {
	for i:= range a{
		fmt.Println(a[i])
	}
}

func main() {
	code_bloat.MakeListAgain()
	fmt.Println(code_bloat.StringList)

	mapIn := make(map[string]a)
	mapIn["i"] = a{}
	mapIn["ii"] = mapIn["ii"]
	for i:=range mapIn  {
		fmt.Println(mapIn[i].Id)
		fmt.Println("name:")
		fmt.Println(mapIn[i].Name)
	}
}
