package bridge_computer

import "fmt"

type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request from WINDOWS")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}
