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

	// gmail
	// create auth method
	gmailServer, gmailAuth := auth.NewAuthGmail("host@address.com", "password", "smtp.gmail.com", ":587")

	// create mail serder & apply auth & call send
	err := mail.NewMailSender(gmailServer).
		Auth(gmailAuth).
		Send(mailer.ToSend())

	// handle error, if need
	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	// office 365
	// create auth method
	// smtp-mail.outlook.com
	// smtp.office365.com
	office365Server, office365Auth := auth.NewAuthOffeice365("host@address.com", "password", "smtp.gmail.com", ":587")

	// create mail serder & apply auth & call send
	err = mail.NewMailSender(office365Server).
		Auth(office365Auth).
		Send(mailer.ToSend())

	// handle error, if need
	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}
