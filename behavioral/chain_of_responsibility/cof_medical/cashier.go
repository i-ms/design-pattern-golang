package cof_medical

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier receiving money from patient")
}

func (c *Cashier) setNext(d Department) {
	c.next = d
}
