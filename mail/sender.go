package mail

import "gopkg.in/gomail.v2"

func SendRegistrationEmail(to, code string, dialer *gomail.Dialer) error {
	m := gomail.NewMessage()

	m.SetHeader("From", dialer.Username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "ASM: verify your email")

	msg := []byte("Dear visitor,\n" +
		"Your email confirmation code is:\n\n" +
		code +
		"\n\nOnce you finish the activation process, you can enter ASM panel and start using our services.")
	m.SetBody("text/plain", string(msg))

	return dialer.DialAndSend(m)
}
