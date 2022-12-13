package template_otp

import "fmt"

func App() {
	smsOTP := &Sms{}
	o := Otp{
		iotp: smsOTP,
	}
	o.genAndSendOTP(4)

	fmt.Println("")
	emailOTP := &Email{}
	o = Otp{
		iotp: emailOTP,
	}
	o.genAndSendOTP(4)
}
