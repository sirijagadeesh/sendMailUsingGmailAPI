package email

import "github.com/sirijagadeesh/sendMailUsingGmailAPI/email/gmailapi"

// Sender will help to send email.
type Sender interface {
	Send(from string, to []string, subject, body string) error
}

// GmailSender will create some.
func GmailSender() (Sender, error) {
	return gmailapi.GmailServiceRepo()
}
