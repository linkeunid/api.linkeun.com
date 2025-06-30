package mails

import (
	"fmt"

	"github.com/goravel/framework/contracts/mail"
	"github.com/goravel/framework/facades"
	"github.com/linkeunid/api.linkeun.com/app/models"
)

type VerifyEmail struct {
	user  *models.User
	token string
}

func NewVerifyEmail(user *models.User, token string) *VerifyEmail {
	return &VerifyEmail{
		user:  user,
		token: token,
	}
}

// Attachments attach files to the mail
func (m *VerifyEmail) Attachments() []string {
	return []string{}
}

// Content set the content of the mail
func (m *VerifyEmail) Content() *mail.Content {
	appUrl := facades.Config().GetString("http.url", "http://localhost:3000")
	verificationUrl := fmt.Sprintf("%s/api/auth/verify/%s", appUrl, m.token)

	return &mail.Content{
		Html: fmt.Sprintf("<h1>Verify Your Email</h1><p>Thanks for registering! Please click the link below to verify your email address:</p><a href=\"%s\">%s</a>", verificationUrl, verificationUrl),
	}
}

// Envelope set the envelope of the mail
func (m *VerifyEmail) Envelope() *mail.Envelope {
	return &mail.Envelope{
		To:      []string{m.user.Email},
		Subject: "Verify Your Email Address",
	}
}

// Queue set the queue of the mail
func (m *VerifyEmail) Queue() *mail.Queue {
	return nil
}
