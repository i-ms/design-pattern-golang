package prototype_payment

import "strconv"

type PayPal struct {
	UserEmail string
}

func (p *PayPal) Pay(amount float64) string {
	return "paid using paypal for amount " + strconv.FormatFloat(amount, 'f', -1, 64)
}

func (p *PayPal) Refund(amount float64) string {
	return "refunded using paypal for amount " + strconv.FormatFloat(amount, 'f', -1, 64)
}

func (p *PayPal) Clone() PaymentMethod {
	return &PayPal{UserEmail: p.UserEmail}
}
