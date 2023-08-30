package decorator_payment

type FraudDetectionDecorator struct {
	PaymentDecorator
}

func (f *FraudDetectionDecorator) ProcessPayment(amount float64) string {
	response := "Checking for Fraud. "
	if amount < 0 {
		response += "Invalid amount. "
	}
	return response + f.Method.ProcessPayment(amount)
}
