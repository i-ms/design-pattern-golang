package abstract_factory_payment

type AccountFactory interface {
	CreatePersonalAccount() PaymentMethod
	CreateBusinessAccount() PaymentMethod
}
