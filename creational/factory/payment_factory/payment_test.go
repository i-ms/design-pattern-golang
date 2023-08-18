package payment_factory

import "testing"

func TestPaymentFactory(t *testing.T) {
	factory := &PaymentFactory[int]{}
	methods := []string{"paypal", "creditcard", "banktransfer"}

	for _, m := range methods {
		payment, _ := factory.GetPaymentMethod(m)
		msg := payment.ProcessPayment(1)
		t.Logf("Payment method %s: %s", m, msg)
	}
}
