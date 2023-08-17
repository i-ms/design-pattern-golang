package payment

import "fmt"

// BankTransfer type
type BankTransfer[T Number] struct{}

func (u *BankTransfer[T]) ProcessPayment(amount T) string {
	return fmt.Sprintf("BankTransfer payment processed for amount %v", amount)
}
