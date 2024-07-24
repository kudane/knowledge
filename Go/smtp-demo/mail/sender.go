package mail

import "net/smtp"

type mailSender struct {
	addr string
	auth smtp.Auth
}

func NewMailSender(addr string) MailSender {
	return &mailSender{
		addr: addr,
		auth: nil,
	}
}

func (m *mailSender) Auth(a smtp.Auth) MailSender {
	m.auth = a
	return m
}

func (m *mailSender) Send(md MailBody) error {
	err := smtp.SendMail(m.addr, m.auth, md.From, md.To, md.msg)

	if err != nil {
		return err
	}

	return nil
}
