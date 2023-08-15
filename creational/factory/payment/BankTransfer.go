package payment

import "fmt"

// BankTransfer type
type BankTransfer struct{}

func (u *BankTransfer) ProcessPayment(amount float64) string {
	return fmt.Sprintf("BankTransfer payment processed for amount %f", amount)
}
