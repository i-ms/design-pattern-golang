package openClosed

import "fmt"

type CreditCardPayment struct {
	Amount       float64
	CardNumber   string
	ExpiryDate   string
	SecurityCode string
}

func (p CreditCardPayment) GetAmount() float64 {
	return p.Amount
}

type CreditCardPaymentProcessor struct {
	// Credit card payment processor implementation
}

func (p *CreditCardPaymentProcessor) ProcessPayment(payment Payment) {
	creditCardPayment := payment.(CreditCardPayment)
	// Perform credit card payment processing logic
	fmt.Println("Processing credit card payment:", creditCardPayment.Amount)
}
