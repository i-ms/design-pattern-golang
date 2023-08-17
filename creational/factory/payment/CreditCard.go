package payment

import "fmt"

// CreditCard type
type CreditCard[T Number] struct{}

func (c *CreditCard[T]) ProcessPayment(amount T) string {
	return fmt.Sprintf("Credit card payment processed for amount %v", amount)
}
