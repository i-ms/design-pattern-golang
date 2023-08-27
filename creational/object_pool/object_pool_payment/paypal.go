package object_pool_payment

import "strconv"

type Paypal struct{}

func (p *Paypal) Pay(amount float64) string {
	return "paid using paypal for amount " + strconv.FormatFloat(amount, 'f', -1, 64)
}

func (p *Paypal) Refund(amount float64) string {
	return "refunded using paypal for amount " + strconv.FormatFloat(amount, 'f', -1, 64)
}
