package adapter_payment

import "testing"

func TestPaymentAdapter(t *testing.T) {
	paypal := &Paypal{}
	paypalAdapter := NewPaypalAdapter(paypal)
	paymentResult := ProcessPayment(paypalAdapter)
	if paymentResult != "Paypal payment authorized." {
		t.Errorf("Expected payment result to be 'Paypal payment authorized.', got '%s'", paymentResult)
	}
}
