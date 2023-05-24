package interfaceSegregation

import "fmt"

type CreditCardPayment struct {
	AmountValue float64
}

func (p CreditCardPayment) Amount() float64 {
	return p.AmountValue
}

type CreditCardProcessor struct{}

func (p *CreditCardProcessor) ProcessPayment(payment Payment) {
	// Perform credit card payment processing logic
	fmt.Println("Processing credit card payment:", payment.Amount())
}

func (p *CreditCardProcessor) ProcessRefund(payment Payment) {
	// Perform credit card refund processing logic
	fmt.Println("Processing credit card refund:", payment.Amount())
}
