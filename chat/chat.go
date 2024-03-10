package chat

import (
	"context"
	"encore.dev/rlog"
)

// Service chat posts notifications to a Slack Webhook.
type Service struct {
	webhookURL string
}

// initService is automatically called by Encore when the service starts up.
func initService() (*Service, error) {
	svc := &Service{
		webhookURL: secrets.SlackWebhookURL,
	}
	return svc, nil
}

// Notify is a private endpoint that sends a Slack message to a pre-configured channel.
// Learn more about Encore's API access controls: https://encore.dev/docs/primitives/services-and-apis#access-controls
//
// encore:api private
func (s *Service) Notify(ctx context.Context, p *NotifyParams) error {
	// If the secret isn't set for local development, log a message
	// and do nothing.
	if s.webhookURL == "" {
		rlog.Debug("no slack webhook url defined, skipping slack notification")
		return nil
	}
}
