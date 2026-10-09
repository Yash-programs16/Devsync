package email

type Sender interface {
	SendEmail(
		to string,
		subject string,
		body string,
	) error
}
