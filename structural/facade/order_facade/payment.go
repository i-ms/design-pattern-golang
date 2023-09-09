package orderfacade

import "fmt"

type PaymentProcessor struct{}

func (pp *PaymentProcessor) ProcessPayment(amount float64) bool {
	fmt.Printf("Processing payment for amount : %f\n", amount)
	return true
}
