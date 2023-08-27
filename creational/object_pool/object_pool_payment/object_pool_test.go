package object_pool_payment

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObjectPool(t *testing.T) {
	pool := &PaymentMethodPool{
		available: []PaymentMethod{&Paypal{}, &Stripe{}},
		inUse:     make([]PaymentMethod, 0),
	}

	assert.Equal(t, 2, len(pool.available), "Expected 2 available payment methods")

	pm := pool.Acquire()
	assert.NotNil(t, pm, "Expected a payment method, got nil")

	assert.Equal(t, 2, len(pool.available)+len(pool.inUse), "Pool size mismatch, expected 2")
}

func TestPaymentMethodPool_Release(t *testing.T) {
	pool := &PaymentMethodPool{
		available: []PaymentMethod{},
		inUse: []PaymentMethod{
			&Paypal{},
			&Stripe{},
		},
	}

	assert.Equal(t, 2, len(pool.inUse), "Expected 2 in-use payment methods")

	pm := &Paypal{}
	pool.Release(pm)

	assert.Equal(t, 3, len(pool.available)+len(pool.inUse), "Pool size mismatch, expected 3")
}
