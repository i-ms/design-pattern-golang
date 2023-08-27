package adapter_payment

type PaypalAdapter struct {
	paypal *Paypal
}

func NewPaypalAdapter(paypal *Paypal) *PaypalAdapter {
	return &PaypalAdapter{paypal: paypal}
}

func (p *PaypalAdapter) Authorize() string {
	return p.paypal.Confirm()
}
