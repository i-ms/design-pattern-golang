package payment

import "fmt"

// CreditCard type
type CreditCard struct{}

func (c *CreditCard) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Credit card payment processed for amount %f", amount)
}
