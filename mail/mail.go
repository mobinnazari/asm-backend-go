package mail

import (
	"log"

	"gopkg.in/gomail.v2"
)

func Init(host, username, password string, port int) *gomail.Dialer {
	dialer := gomail.NewDialer(host, port, username, password)
	closer, err := dialer.Dial()
	if err != nil {
		log.Panicln(err.Error())
	}

	closer.Close()
	return dialer
}
