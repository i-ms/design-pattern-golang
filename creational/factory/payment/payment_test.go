package payment

import "testing"

func TestPaymentFactory(t *testing.T) {
	factory := &PaymentFactory{}
	methods := []string{"paypal", "creditcard", "banktransfer"}

	for _, m := range methods {
		payment, _ := factory.GetPaymentMethod(m)
		msg := payment.ProcessPayment(10.0)
		t.Logf("Payment method %s: %s", m, msg)
	}
}
