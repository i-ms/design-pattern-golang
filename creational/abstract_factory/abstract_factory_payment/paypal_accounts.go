package abstract_factory_payment

import "fmt"

type PaypalPersonalAccount struct{}
type PaypalBusinessAccount struct{}

func (p *PaypalPersonalAccount) Pay(amount float64) string {
	return fmt.Sprintf("PaypalPersonalAccount.Pay : %f", amount)
}

func (p *PaypalPersonalAccount) Refund(amount float64) string {
	return fmt.Sprintf("PaypalPersonalAccount.Refund : %f", amount)
}

func (p *PaypalPersonalAccount) Transfer(to PaymentMethod, amount float64) string {
	return fmt.Sprintf("PaypalPersonalAccount.Transfer : %f", amount)
}

func (p *PaypalBusinessAccount) Pay(amount float64) string {
	return fmt.Sprintf("PaypalBusinessAccount.Pay : %f", amount)
}

func (p *PaypalBusinessAccount) Refund(amount float64) string {
	return fmt.Sprintf("PaypalBusinessAccount.Refund : %f", amount)
}

func (p *PaypalBusinessAccount) Transfer(to PaymentMethod, amount float64) string {
	return fmt.Sprintf("PaypalBusinessAccount.Transfer : %f", amount)
}
