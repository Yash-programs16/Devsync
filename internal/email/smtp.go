package email

import (
	"fmt"
	"net/smtp"
)

type SMTPClient struct {
	host string

	port string

	username string

	password string
}

func NewSMTPClient(
	host string,
	port string,
	username string,
	password string,
) *SMTPClient {

	return &SMTPClient{

		host: host,

		port: port,

		username: username,

		password: password,
	}
}
func (s *SMTPClient) SendEmail(
	to string,
	subject string,
	body string,
) error {

	message := []byte(
		fmt.Sprintf(
			"Subject: %s\r\n\r\n%s",
			subject,
			body,
		),
	)

	auth := smtp.PlainAuth(
		"",
		s.username,
		s.password,
		s.host,
	)

	err := smtp.SendMail(
		s.host+":"+s.port,
		auth,
		s.username,
		[]string{to},
		message,
	)

	return err
}
