package decorator_payment

type CurrencyConversionDecorator struct {
	PaymentDecorator
}

func (c *CurrencyConversionDecorator) ProcessPayment(amount float64) string {
	response := "Converting currency. "
	convertedAmount := amount * 1.8
	return response + c.Method.ProcessPayment(convertedAmount)
}
