package bridge_computer

import "fmt"

type Epson struct{}

func (e *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON printer")
}
