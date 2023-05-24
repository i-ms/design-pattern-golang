package liskovSubstitution

import "fmt"

type CreditCardPayment struct {
	CreditCardNumber string
	ExpiryDate       string
	CVV              int
}

func (c *CreditCardPayment) Process(amount float32) error {
	// Process the payment
	fmt.Println("Processing credit card payment for amount:", amount)
	return nil
}
