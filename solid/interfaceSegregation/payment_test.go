package interfaceSegregation

import (
	"fmt"
	"testing"
)

func TestPayment(t *testing.T) {
	creditCardPayment := &CreditCardPayment{
		AmountValue: 100,
	}

	creditCardProcessor := &CreditCardProcessor{}

	fmt.Println("Payment processing:")
	creditCardProcessor.ProcessPayment(creditCardPayment)

	fmt.Println("Refund processing:")
	creditCardProcessor.ProcessRefund(creditCardPayment)
}
