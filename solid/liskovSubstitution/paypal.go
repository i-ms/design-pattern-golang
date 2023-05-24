package liskovSubstitution

import "fmt"

type PaypalPayment struct {
	Username string
	Password string
}

func (p *PaypalPayment) Process(amount float32) error {
	// Process the payment
	fmt.Println("Processing paypal payment for amount:", amount)
	return nil
}
