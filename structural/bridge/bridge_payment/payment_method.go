package bridge_payment

type PaymentMethod interface {
	Process() string
	Cancel() string
	GetDetails() string
}
