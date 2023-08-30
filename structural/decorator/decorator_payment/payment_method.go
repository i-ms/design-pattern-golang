package decorator_payment

type PaymentMethod interface {
	ProcessPayment(float64) string
}
