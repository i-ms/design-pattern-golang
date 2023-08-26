package prototype_payment

import "strconv"

type Stripe struct {
	AccountID string
}

func (s *Stripe) Pay(amount float64) string {
	return "paid using stripe for amount " + strconv.FormatFloat(amount, 'f', -1, 64)
}

func (s *Stripe) Refund(amount float64) string {
	return "refunded using stripe for amount " + strconv.FormatFloat(amount, 'f', -1, 64)
}

func (s *Stripe) Clone() PaymentMethod {
	return &Stripe{AccountID: s.AccountID}
}
