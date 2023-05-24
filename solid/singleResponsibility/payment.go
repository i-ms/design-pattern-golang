package singleResponsibility

type Payment interface {
	Process() error
}

type PaymentProcessor struct{}

func (p *PaymentProcessor) Process(payment Payment) error {
	return payment.Process()
}
