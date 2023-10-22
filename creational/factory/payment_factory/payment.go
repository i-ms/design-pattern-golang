package payment_factory

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}
type PaymentProcessor[T Number] interface {
	ProcessPayment(amount T) string
}

type PaymentFactory[T Number] struct{}

func (p *PaymentFactory[T]) GetPaymentMethod(method string) (PaymentProcessor[T], error) {
	switch method {
	case "paypal":
		return new(Paypal[T]), nil
	case "creditcard":
		return new(CreditCard[T]), nil
	case "banktransfer":
		return new(BankTransfer[T]), nil
	default:
		return nil, fmt.Errorf("Payment method %s not recognized", method)
	}
}
