package mails

import (
	"github.com/goravel/framework/contracts/mail"
	"github.com/linkeunid/api.linkeun.com/app/models"
)

type UserRegistered struct {
	user *models.User
}

func NewUserRegistered(user *models.User) *UserRegistered {
	return &UserRegistered{
		user: user,
	}
}

// Attachments attach files to the mail
func (m *UserRegistered) Attachments() []string {
	return []string{}
}

// Content set the content of the mail
func (m *UserRegistered) Content() *mail.Content {
	return &mail.Content{
		Html: "<h1>Welcome " + m.user.Name + "</h1><p>Thank you for registering!</p>",
	}
}

// Envelope set the envelope of the mail
func (m *UserRegistered) Envelope() *mail.Envelope {
	return &mail.Envelope{
		To:      []string{m.user.Email},
		Subject: "Welcome to Our Application!",
	}
}

// Queue set the queue of the mail
func (m *UserRegistered) Queue() *mail.Queue {
	return &mail.Queue{
		Connection: "redis",
		Queue:      "mails",
	}
}
