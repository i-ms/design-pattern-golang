package payment

import "fmt"

// Paypal type
type Paypal struct{}

func (p *Paypal) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Paypal payment processed for amount %f", amount)
}
