package main

import "fmt"

// IProcess Interface
type IProcess interface {
	process()
}

// ProcessClass struct
type ProcessClass struct{}

// ProcessClass method process
func (process *ProcessClass) process() {
	fmt.Println("ProcessClass process")
}

// ProcessDecorator struct
type ProcessDecorator struct {
	processInstance *ProcessClass
}

// ProcessDecorator class method process
func (processDecorator *ProcessDecorator) process() {
	if processDecorator.processInstance == nil {
		fmt.Println("ProcessDecorator process")
	} else {
		fmt.Printf("ProcessDecorator process and ")
		processDecorator.processInstance.process()
	}
}

func main() {
	var (
		process   = &ProcessClass{}
		decorator = &ProcessDecorator{}
	)
	decorator.process()
	decorator.processInstance = process
	decorator.process()
}
