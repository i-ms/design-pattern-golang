package liskovSubstitution

import "fmt"

type DebitCardPayment struct {
	DebitCardNumber string
	ExpiryDate      string
	CVV             int
}

func (d *DebitCardPayment) Process(amount float32) error {
	// Process the payment
	fmt.Println("Processing debit card payment for amount:", amount)
	return nil
}
