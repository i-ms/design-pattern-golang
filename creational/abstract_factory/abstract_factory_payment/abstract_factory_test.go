package abstract_factory_payment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaypalPersonalAccount_Pay(t *testing.T) {
	ppa := PaypalPersonalAccount{}
	result := ppa.Pay(100.0)
	assert.Equal(t, "PaypalPersonalAccount.Pay : 100.000000", result)
}

func TestPaypalBusinessAccount_Pay(t *testing.T) {
	pba := PaypalBusinessAccount{}
	result := pba.Pay(200.0)
	assert.Equal(t, "PaypalBusinessAccount.Pay : 200.000000", result)
}

func TestStripePersonalAccount_Pay(t *testing.T) {
	spa := StripePersonalAccount{}
	result := spa.Pay(300.0)
	assert.Equal(t, "StripePersonalAccount.Pay : 300.000000", result)
}

func TestStripeBusinessAccount_Pay(t *testing.T) {
	sba := StripeBusinessAccount{}
	result := sba.Pay(400.0)
	assert.Equal(t, "StripeBusinessAccount.Pay : 400.000000", result)
}
