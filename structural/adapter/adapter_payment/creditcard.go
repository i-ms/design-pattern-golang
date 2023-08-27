package adapter_payment

type CreditCard struct{}

func (c *CreditCard) Validate() string {
	return "Credit card payment Validated."
}
