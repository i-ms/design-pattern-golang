package decorator_payment

type LoggingDecorator struct {
	PaymentDecorator
}

func (l *LoggingDecorator) ProcessPayment(amount float64) string {
	response := "Logging payment. "
	return response + l.Method.ProcessPayment(amount)
}
