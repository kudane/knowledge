package auth

import mail "net/smtp"

func NewAuthGmail(username, password, host string) mail.Auth {
	return mail.PlainAuth("", username, password, host)
}
