package template_otp

type Otp struct {
	iotp Iotp
}

func (o *Otp) genAndSendOTP(otpLength int) error {
	otp := o.iotp.genRandomOTP(otpLength)
	o.iotp.saveOTPCache(otp)
	message := o.iotp.getMessage(otp)
	err := o.iotp.sendNotification(message)
	if err != nil {
		return err
	}
	return nil
}
