package mail

import "gopkg.in/gomail.v2"

func CheckHealth(dialer *gomail.Dialer) error {
	closer, err := dialer.Dial()
	if err != nil {
		return err
	}

	closer.Close()
	return nil
}
