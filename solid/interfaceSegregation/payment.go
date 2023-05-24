package interfaceSegregation

type Payment interface {
	Amount() float64
}

type PaymentProcessor interface {
	ProcessPayment(payment Payment)
}

type RefundProcessor interface {
	ProcessRefund(payment Payment)
}

type PaymentRefundProcessor interface {
	PaymentProcessor
	RefundProcessor
}
