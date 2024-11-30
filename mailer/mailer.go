package mailer

type Mailer interface {
	Send(to, body string) error
	ReadTemplate(string, any) (string, error)
}
