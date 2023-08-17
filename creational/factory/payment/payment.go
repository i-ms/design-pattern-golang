package payment

import "fmt"

type PaymentProcessor interface {
	ProcessPayment(amount float64) string
}

type PaymentFactory struct{}

func (p *PaymentFactory) GetPaymentMethod(method string) (PaymentProcessor, error) {
	switch method {
	case "paypal":
		return new(Paypal), nil
	case "creditcard":
		return new(CreditCard), nil
	case "banktransfer":
		return new(BankTransfer), nil
	default:
		return nil, fmt.Errorf("Payment method %s not recognized", method)
	}
}
