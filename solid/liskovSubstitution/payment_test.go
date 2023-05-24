package liskovSubstitution

import "testing"

func TestPayment(t *testing.T) {
	creditCard := &CreditCardPayment{
		CreditCardNumber: "1234 1234 1234 1234",
		ExpiryDate:       "12/21",
		CVV:              123,
	}
	_ = ProcessPayment(creditCard, 10)

	debitCard := &DebitCardPayment{
		DebitCardNumber: "1234 1234 1234 1234",
		ExpiryDate:      "12/21",
		CVV:             123,
	}
	_ = ProcessPayment(debitCard, 200)

	paypal := &PaypalPayment{
		Username: "sampleUserName",
		Password: "password",
	}
	_ = ProcessPayment(paypal, 500)
}
