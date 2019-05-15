package main

import "fmt"

type WithOut struct {
	MainID int
	Id int
}

func (w WithOut) printOk(it string) {
	fmt.Println("With out ", it)
}

func (w WithOut) printSelf(it string) {
	fmt.Println(w.Id, it)
	fmt.Println(w.MainID, it)
}

type WithIn struct {
	Id int
}

func (w WithIn) printOk(it string) {
	fmt.Println("With in ", it)
}

func (w WithIn) printSelf(it string) {
	fmt.Println(w.Id, it)
}

type nameInterface interface {
	printOk(string)
	printSelf(string)
}

func main() {
	var test nameInterface
	if false {
		test = WithIn{Id: 10}
		test.printOk("it")
		test.printSelf("it")
	} else {
		test = WithOut{Id: 10, MainID:20}
		test.printOk("it")
		test.printSelf("it")
	}
}
