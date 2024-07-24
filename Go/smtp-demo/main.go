package main

import (
	"github/kudane/smtp-demo/mail"
	"github/kudane/smtp-demo/mail/auth"
	"log"
)

func main() {
	// defind mail body
	from := "yourmail@example.com"
	to := "to@example.com"
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n"

	// create mail data
	mailer := mail.NewMailMessage().
		From(from).
		To(to).
		Body([]byte(msg))

	// create auth method
	auth := auth.NewAuthGmail("host@address.com", "password", "smtp.gmail.com")

	// create mail serder & apply auth & call send
	err := mail.NewMailSender("smtp.gmail.com:587").
		Auth(auth).
		Send(mailer.ToSend())

	// handle error, if need
	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}
