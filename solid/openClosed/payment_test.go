package openClosed

import "testing"

func TestPaymentProcessors(t *testing.T) {
	creditCardPayment := CreditCardPayment{
		Amount:       100.0,
		CardNumber:   "1234567890",
		ExpiryDate:   "12/25",
		SecurityCode: "123",
	}

	paypalPayment := PayPalPayment{
		Amount:   200.0,
		Email:    "user@example.com",
		Password: "password",
	}

	bankTransferPayment := BankTransferPayment{
		Amount:      300.0,
		Account:     "9876543210",
		RoutingCode: "123456",
	}

	creditCardProcessor := &CreditCardPaymentProcessor{}
	paypalProcessor := &PayPalPaymentProcessor{}
	bankTransferProcessor := &BankTransferPaymentProcessor{}

	creditCardProcessor.ProcessPayment(creditCardPayment)
	paypalProcessor.ProcessPayment(paypalPayment)
	bankTransferProcessor.ProcessPayment(bankTransferPayment)

}
