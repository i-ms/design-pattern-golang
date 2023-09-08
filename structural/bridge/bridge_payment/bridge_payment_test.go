package bridge_payment

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBridgePayment(t *testing.T) {
	assertion := assert.New(t)
	onlinePayment := &OnlinePayment{}
	onlinePayment.SetPaymentMethod(&CreditCard{})
	// Test success scenario
	assertion.Equal("Online Payment: Processed by Credit Card.", onlinePayment.Authorize(), "OnlinePayment with CreditCard should succeed.")
	assertion.Equal("Online Payment Refund: Transaction Cancelled by Credit Card.", onlinePayment.Refund(), "OnlinePayment Refund with CreditCard should succeed.")
	assertion.Equal("Online Payment Status: Transaction Details from Credit Card.", onlinePayment.GetStatus(), "OnlinePayment GetStatus with CreditCard should succeed.")

	// Test failure scenario
	onlinePayment.SetPaymentMethod(nil)
	assertion.NotEqual("Online Payment: Processed by Credit Card.", onlinePayment.Authorize(), "OnlinePayment with nil PaymentMethod should fail.")
}
