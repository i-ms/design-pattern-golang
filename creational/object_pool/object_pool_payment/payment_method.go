package object_pool_payment

type PaymentMethod interface {
	Pay(amount float64) string
	Refund(amount float64) string
}
