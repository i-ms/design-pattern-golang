package singleResponsibility

import "testing"

func TestPayment(t *testing.T) {
	creditCardPayment := &CreditCardPayment{
		Amount:           149,
		CustomerEmail:    "credit@example.com",
		CreditCardNumber: "1234 4562 3214",
		ExpiryDate:       "24/02",
		CVV:              259,
	}

	paypalPayment := &PaypalPayment{
		Amount:        150,
		CustomerEmail: "credit@example.com",
		Username:      "sample",
		Password:      "password",
	}

	paymentProcessor := &PaymentProcessor{}
	_ = paymentProcessor.Process(creditCardPayment)
	_ = paymentProcessor.Process(paypalPayment)
}
