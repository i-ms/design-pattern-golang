package openClosed

import "fmt"

type BankTransferPayment struct {
	Amount      float64
	Account     string
	RoutingCode string
}

func (p BankTransferPayment) GetAmount() float64 {
	return p.Amount
}

type BankTransferPaymentProcessor struct {
	// Bank transfer payment processor implementation
}

func (p *BankTransferPaymentProcessor) ProcessPayment(payment Payment) {
	bankTransferPayment := payment.(BankTransferPayment)
	// Perform bank transfer payment processing logic
	fmt.Println("Processing bank transfer payment:", bankTransferPayment.Amount)
}
