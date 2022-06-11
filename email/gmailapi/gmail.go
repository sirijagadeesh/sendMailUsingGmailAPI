package gmailapi

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"strings"
	"time"

	"github.com/sirijagadeesh/sendMailUsingGmailAPI/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

// Mail generic structure.
type Mail struct {
	From       string   `json:"from,omitempty" validate:"required,email"`
	To         []string `json:"to,omitempty" validate:"required,dive,email"`
	Subject    string   `json:"subject,omitempty" validate:"required"`
	Body       string   `json:"body,omitempty" validate:"required"`
	Template   string   `json:"template" validate:"required,oneof=simple1 simple2 simple3"`
	Recipients string   `json:"recipient_mailList,omitempty" validate:"required"`
}

// Message will get gmail.Message we need.
func (m Mail) Message() (*gmail.Message, error) {
	emailTemplate, err := template.ParseFS(config.Templates, fmt.Sprintf("templates/%s.tmpl", m.Template))
	if err != nil {
		return nil, fmt.Errorf("unable to parse template: %w", err)
	}

	var messageBody bytes.Buffer

	if err := emailTemplate.Execute(&messageBody, m); err != nil {
		return nil, fmt.Errorf("unable execute template %w", err)
	}

	var message gmail.Message
	message.Raw = base64.URLEncoding.EncodeToString(messageBody.Bytes())

	return &message, nil
}

// GmailService which will send all details.
type GmailService struct {
	*gmail.Service
}

// GmailServiceRepo get gmail service.
func GmailServiceRepo() (*GmailService, error) {
	ctx := context.Background()
	apiConfig := config.GmailAPI()

	config := &oauth2.Config{
		ClientID:     apiConfig.GmailClientID,
		ClientSecret: apiConfig.GmailClientSecret,
		RedirectURL:  apiConfig.GmailAPIRedirectURI,
		Endpoint:     google.Endpoint,
		Scopes:       nil,
	}

	accessToken := config.TokenSource(ctx, &oauth2.Token{
		AccessToken:  "",
		RefreshToken: apiConfig.GmailAPIRefreshToken,
		TokenType:    "",
		Expiry:       time.Now(),
	})

	gmailService, err := gmail.NewService(ctx, option.WithTokenSource(accessToken))
	if err != nil {
		return nil, fmt.Errorf("unable to get gmail service %w", err)
	}

	return &GmailService{gmailService}, nil
}

// Send will send mail using GMAIL API.
func (gs *GmailService) Send(from string, to []string, subject, body string) error {
	email := Mail{from, to, subject, body, "simple1", strings.Join(to, ",")}

	message, err := email.Message()
	if err != nil {
		return fmt.Errorf("unable to crete e-mail %w", err)
	}

	msg, err := gs.Users.Messages.Send("me", message).Do()
	if err != nil {
		return fmt.Errorf("unable to send e-mail : %w", err)
	}

	log.Println(msg)

	return nil
}
