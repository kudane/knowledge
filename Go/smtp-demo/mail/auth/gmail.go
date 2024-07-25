package auth

import "net/smtp"

func NewAuthGmail(username, password, host, port string) (string, smtp.Auth) {
	server := host + port
	auth := smtp.PlainAuth("", username, password, host)
	return server, auth
}
