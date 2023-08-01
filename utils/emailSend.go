package utils

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

type EmailSend interface {
	SendEmail(
		subject string,
		content string,
		to []string,
		cc []string,
		bcc []string,
	) error
}

type GmailSend struct {
	name          string
	EmailAddress  string
	EmailPassword string
}

func NewGmailSend(name string, EmailAddress string, EmailPassword string) EmailSend {
	return &GmailSend{
		name:          name,
		EmailAddress:  EmailAddress,
		EmailPassword: EmailPassword,
	}
}

func (sender *GmailSend) SendEmail(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
) error {
	e := email.NewEmail()

	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.EmailAddress)
	e.Subject = subject
	// content string, html is []byte
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	//smtp auth
	// first is empty,hour is smtp auth server(in case gmail)
	smtpAuth := smtp.PlainAuth("", sender.EmailAddress, sender.EmailPassword, smtpAuthAddress)
	return e.Send(smtpServerAddress, smtpAuth)

}
