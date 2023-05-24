package liskovSubstitution

type Payment interface {
	Process(amount float32) error
}

func ProcessPayment(p Payment, amount float32) error {
	return p.Process(amount)
}
