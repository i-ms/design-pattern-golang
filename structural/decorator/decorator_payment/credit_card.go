package decorator_payment

import "fmt"

type CreditCard struct{}

func (c *CreditCard) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Processed payment of %f through Credit Card", amount)
}
