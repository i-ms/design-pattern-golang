package payment

import "fmt"

type PaymentProcessor interface {
	ProcessPayment(amount float64) string
}

// CreditCard type
type CreditCard struct{}

func (c *CreditCard) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Credit card payment processed for amount %f", amount)
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
