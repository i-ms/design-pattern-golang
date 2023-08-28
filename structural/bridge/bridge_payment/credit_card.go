package bridge_payment

type CreditCard struct{}

func (c *CreditCard) Process() string {
	return "Processed by Credit Card."
}

func (c *CreditCard) Cancel() string {
	return "Transaction Cancelled by Credit Card."
}

func (c *CreditCard) GetDetails() string {
	return "Transaction Details from Credit Card."
}

// More code for other online payment types can be added like PayPal, stripe, UPI etc.
