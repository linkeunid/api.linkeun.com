package mails

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/goravel/framework/contracts/mail"
	"github.com/goravel/framework/facades"
	"github.com/linkeunid/api.linkeun.com/app/models"
)

//go:embed templates/email_verify.html
var emailTemplate string

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
	appUrl := facades.Config().GetString("http.frontend_url", "http://localhost:5173")

	verificationUrl := fmt.Sprintf("%s/auth/verify/%s", appUrl, m.token)

	htmlContent, err := m.generateVerificationEmail(verificationUrl)
	if err != nil {
		// Fallback to simple HTML if template fails
		htmlContent = fmt.Sprintf("<h1>Verify Your Email</h1><p>Thanks for registering! Please click the link below to verify your email address:</p><a href=\"%s\">%s</a>", verificationUrl, verificationUrl)
	}

	return &mail.Content{
		Html: htmlContent,
	}
}

// generateVerificationEmail creates the HTML content using the email template
func (m *VerifyEmail) generateVerificationEmail(verificationUrl string) (string, error) {
	// Get configuration values
	companyName := facades.Config().GetString("app.name", "LinkeUnID")
	supportEmail := facades.Config().GetString("mail.from.support.address")
	currentYear := fmt.Sprintf("%d", time.Now().Year())

	// Replace placeholders
	replacer := strings.NewReplacer(
		"{{VERIFICATION_URL}}", verificationUrl,
		"{{COMPANY_NAME}}", companyName,
		"{{SUPPORT_EMAIL}}", supportEmail,
		"{{CURRENT_YEAR}}", currentYear,
	)

	return replacer.Replace(emailTemplate), nil
}

// Envelope set the envelope of the mail
func (m *VerifyEmail) Envelope() *mail.Envelope {
	companyName := facades.Config().GetString("app.name", "LinkeUnID")

	return &mail.Envelope{
		To:      []string{m.user.Email},
		Subject: fmt.Sprintf("%s - Verify Your Email Address", companyName),
	}
}

// Queue set the queue of the mail
func (m *VerifyEmail) Queue() *mail.Queue {
	return nil
}
