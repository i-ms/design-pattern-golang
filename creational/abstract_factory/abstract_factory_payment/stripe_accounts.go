package abstract_factory_payment

import "fmt"

type StripePersonalAccount struct{}
type StripeBusinessAccount struct{}

func (s *StripePersonalAccount) Pay(amount float64) string {
	return fmt.Sprintf("StripePersonalAccount.Pay : %f", amount)
}

func (s *StripePersonalAccount) Refund(amount float64) string {
	return fmt.Sprintf("StripePersonalAccount.Refund : %f", amount)
}

func (s *StripePersonalAccount) Transfer(to PaymentMethod, amount float64) string {
	return fmt.Sprintf("StripePersonalAccount.Transfer : %f", amount)
}

func (s *StripeBusinessAccount) Pay(amount float64) string {
	return fmt.Sprintf("StripeBusinessAccount.Pay : %f", amount)
}

func (s *StripeBusinessAccount) Refund(amount float64) string {
	return fmt.Sprintf("StripeBusinessAccount.Refund : %f", amount)
}

func (s *StripeBusinessAccount) Transfer(to PaymentMethod, amount float64) string {
	return fmt.Sprintf("StripeBusinessAccount.Transfer : %f", amount)
}
