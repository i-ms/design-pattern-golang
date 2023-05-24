package openClosed

import "fmt"

type PayPalPayment struct {
	Amount   float64
	Email    string
	Password string
}

func (p PayPalPayment) GetAmount() float64 {
	return p.Amount
}

type PayPalPaymentProcessor struct {
	// PayPal payment processor implementation
}

func (p *PayPalPaymentProcessor) ProcessPayment(payment Payment) {
	payPalPayment := payment.(PayPalPayment)
	// Perform PayPal payment processing logic
	fmt.Println("Processing PayPal payment:", payPalPayment.Amount)
}
