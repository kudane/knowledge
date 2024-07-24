package mail

import "net/smtp"

type MailMessage interface {
	From(f string) MailMessage
	To(t string) MailMessage
	Body(b []byte) MailMessage
	ToSend() MailBody
}

type MailSender interface {
	Auth(a smtp.Auth) MailSender
	Send(m MailBody) error
}

type MailBody struct {
	From string
	To   []string
	msg  []byte
}
