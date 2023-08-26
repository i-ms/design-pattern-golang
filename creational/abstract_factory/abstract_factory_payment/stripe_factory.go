package abstract_factory_payment

type StripeAccountFactory struct{}

func (s *StripeAccountFactory) CreatePersonalAccount() PaymentMethod {
	return &StripePersonalAccount{}
}

func (s *StripeAccountFactory) CreateBusinessAccount() PaymentMethod {
	return &StripeBusinessAccount{}
}
