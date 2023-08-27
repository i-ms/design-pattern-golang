package adapter_payment

type CreditCardAdapter struct {
	cc *CreditCard
}

func NewCreditCardAdapter(cc *CreditCard) *CreditCardAdapter {
	return &CreditCardAdapter{cc: cc}
}

func (c *CreditCardAdapter) Authorize() string {
	return c.cc.Validate()
}
