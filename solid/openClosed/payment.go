package openClosed

type Payment interface {
	GetAmount() float64
}

type PaymentProcessor interface {
	ProcessPayment(payment Payment)
}
