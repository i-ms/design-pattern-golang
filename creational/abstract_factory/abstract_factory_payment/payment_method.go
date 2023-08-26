package abstract_factory_payment

// PaymentMethod provide generic interface for account operations
type PaymentMethod interface {
	Pay(amount float64) string
	Refund(amount float64) string
	Transfer(to PaymentMethod, amount float64) string
}
