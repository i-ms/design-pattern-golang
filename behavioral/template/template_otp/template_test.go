package template_otp

import (
	"fmt"
	"testing"
)

func TestTemplate(t *testing.T) {
	smsOTP := &Sms{}
	o := Otp{
		iotp: smsOTP,
	}
	_ = o.genAndSendOTP(4)

	fmt.Println("")
	emailOTP := &Email{}
	o = Otp{
		iotp: emailOTP,
	}
	_ = o.genAndSendOTP(4)
}
