package email

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mrz1836/postmark"
)

type sender struct {
	email     string
	templates map[string]EmailTemplate
	client    *postmark.Client
}

type EmailTemplate struct {
	ID     int64                  `json:"id"`
	Values map[string]interface{} `json:"values"`
}

func NewSender(serverToken, fromEmail, emailTemplateDir string) (sender, error) {
	s := sender{
		email:  fromEmail,
		client: postmark.NewClient(serverToken, ""),
	}

	f, err := os.ReadFile(filepath.Clean(emailTemplateDir))
	if err != nil {
		return sender{}, fmt.Errorf("error opening email template directory: %s", err.Error())
	}

	if err := json.Unmarshal(f, &s.templates); err != nil {
		return sender{}, fmt.Errorf("error unmarshalling email templates: %s", err.Error())
	}

	return s, nil
}

func (s sender) SendNotification(ctx context.Context, notificationType, toEmail string, templateBody map[string]interface{}) error {
	t, ok := s.templates[notificationType]
	if !ok {
		return fmt.Errorf("error sending email: invalid notification type %s", notificationType)
	}

	for k, v := range templateBody {
		t.Values[k] = v
	}

	_, err := s.client.SendTemplatedEmail(ctx, postmark.TemplatedEmail{
		TemplateID:    t.ID,
		TemplateModel: t.Values,
		From:          s.email,
		To:            toEmail,
	})
	if err != nil {
		return fmt.Errorf("error sending email: %s", err.Error())
	}

	return nil
}
