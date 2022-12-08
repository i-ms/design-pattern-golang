package bridge_computer

import "fmt"

type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request from MAC")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}
