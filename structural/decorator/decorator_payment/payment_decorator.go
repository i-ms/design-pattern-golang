package decorator_payment

type PaymentDecorator struct {
	Method PaymentMethod
}

func (p *PaymentDecorator) SetMethod(method PaymentMethod) {
	p.Method = method
}

func (p *PaymentDecorator) ProcessPayment(amount float64) string {
	return p.Method.ProcessPayment(amount)
}
