package prototype_payment

import "testing"

func TestPrototype(t *testing.T) {
	stripe := &Stripe{AccountID: "stripe_account_id"}
	paypal := &PayPal{UserEmail: "paypal_user_email"}

	stripeClone := stripe.Clone()
	paypalClone := paypal.Clone()

	if stripeClone == stripe {
		t.Error("Stripe and its clone should not be the same instance")
	}

	if paypalClone == paypal {
		t.Error("PayPal and its clone should not be the same instance")
	}

	println(stripeClone.Pay(100))
	println(paypalClone.Pay(200))
}
