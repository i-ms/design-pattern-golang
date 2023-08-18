package payment_factory

import "fmt"

// Paypal type
type Paypal[T Number] struct{}

func (p *Paypal[T]) ProcessPayment(amount T) string {
	return fmt.Sprintf("Paypal payment processed for amount %v", amount)
}
