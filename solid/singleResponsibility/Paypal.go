package singleResponsibility

import "fmt"

type PaypalPayment struct {
	Amount        float64
	CustomerEmail string
	Username      string
	Password      string
}

func (p *PaypalPayment) Process() error {
	fmt.Printf("Processing Paypapl payment of %f for %s", p.Amount, p.Username)
	return nil
}
