package singleResponsibility

import "fmt"

type CreditCardPayment struct {
	Amount           float64
	CustomerEmail    string
	CreditCardNumber string
	ExpiryDate       string
	CVV              int
}

func (c *CreditCardPayment) Process() error {
	fmt.Printf("Processing Credit Card payment of %f for %s\n", c.Amount, c.CreditCardNumber)
	return nil
}
