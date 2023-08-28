package bridge_payment

type Payment interface {
	Authorize() string
	Refund() string
	GetStatus() string
	SetPaymentMethod(paymentMethod PaymentMethod)
}
