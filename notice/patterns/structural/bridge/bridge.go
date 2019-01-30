package bridge

import (
	"fmt"
	"github.com/kataras/iris/core/errors"
	"io"
)

/*
A PrinterAPI that accepts a message to print
An implementation of the API that simply prints the message to the console
An implementation of the API that prints to an io.Writer interface
A PrinterAPI abstraction with a Print method to implement in printing types
A normal printer object,which will implement the PrinterAPI and the PrinterAPI interface
The normal printer will forward the message directly to the implementation
A Packet printer,which will implement the PrinterAPI abstraction and the PrinterAPI interface
The Packet printer will append the message Message from Packt: to all prints.
*/

type PrinterAPI interface {
	PrintMessage(string) error
}

type PrinterAPI1 struct{}

func (p *PrinterAPI1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg)
	return nil
}

type PrinterAPI2 struct {
	Writer io.Writer
}

func (p *PrinterAPI2) PrintMessage(msg string) error {
	if p.Writer == nil {
		return errors.New("You need to pass an io.Writer to PrinterAPI2")
	}
	_, err := fmt.Fprintf(p.Writer, "%s", msg)
	return err
}

type PrinterAbstraction interface {
	Print() error
}

type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (n *NormalPrinter) Print() error {
	return n.Printer.PrintMessage(n.Msg)
}

type PacktPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (p *PacktPrinter) Print() error {
	return p.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %s", p.Msg))
}
