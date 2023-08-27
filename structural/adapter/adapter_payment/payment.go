package adapter_payment

type Payment interface {
	Authorize() string
}

func ProcessPayment(p Payment) string {
	return p.Authorize()
}
