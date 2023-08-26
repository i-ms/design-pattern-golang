package abstract_factory_payment

type PaypalAccountFactory struct{}

func (p *PaypalAccountFactory) CreatePersonalAccount() PaymentMethod {
	return &PaypalPersonalAccount{}
}

func (p *PaypalAccountFactory) CreateBusinessAccount() PaymentMethod {
	return &PaypalBusinessAccount{}
}
