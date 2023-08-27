package adapter_payment

type Paypal struct{}

func (p *Paypal) Confirm() string {
	return "Paypal payment authorized."
}
