package bridge_payment

type OnlinePayment struct {
	method PaymentMethod
}

func (o *OnlinePayment) Authorize() string {
	if o.method == nil {
		return "Payment method not set"
	}
	return "Online Payment: " + o.method.Process()
}

func (o *OnlinePayment) Refund() string {
	return "Online Payment Refund: " + o.method.Cancel()
}

func (o *OnlinePayment) GetStatus() string {
	return "Online Payment Status: " + o.method.GetDetails()
}

func (o *OnlinePayment) SetPaymentMethod(paymentMethod PaymentMethod) {
	o.method = paymentMethod
}

// Similarly more code for offline payment can be added in new file offline_payment.go
