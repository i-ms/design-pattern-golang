package decorator_payment

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaymentLoggingDecorator(t *testing.T) {
	assert := assert.New(t)

	creditCard := &CreditCard{}
	loggingDecorator := &LoggingDecorator{}
	loggingDecorator.SetMethod(creditCard)

	assert.Equal("Logging payment. Processed payment of 100.000000 through Credit Card", loggingDecorator.ProcessPayment(100), "Payment should be processed and logged.")
}

func TestPaymentFraudDetectionDecorator(t *testing.T) {
	assertion := assert.New(t)
	creditCard := &CreditCard{}
	fraudDetectionDecorator := &FraudDetectionDecorator{}
	fraudDetectionDecorator.SetMethod(creditCard)

	assertion.Equal("Checking for Fraud. Invalid amount. Processed payment of -100.000000 through Credit Card", fraudDetectionDecorator.ProcessPayment(-100), "Payment should be processed and fraud should be detected.")
}

func TestPaymentCurrencyConversionDecorator(t *testing.T) {
	assertion := assert.New(t)
	creditCard := &CreditCard{}
	currencyConversionDecorator := &CurrencyConversionDecorator{}
	currencyConversionDecorator.SetMethod(creditCard)

	assertion.Equal("Converting currency. Processed payment of 180.000000 through Credit Card", currencyConversionDecorator.ProcessPayment(100), "Payment should be processed and converted.")
}
