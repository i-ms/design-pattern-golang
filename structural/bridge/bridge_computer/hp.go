package bridge_computer

import "fmt"

type HP struct{}

func (h *HP) PrintFile() {
	fmt.Println("Printing from HP printer")
}
