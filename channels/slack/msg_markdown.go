package slack

import (
	"github.com/bougou/webhook-adapter/models"
	"github.com/slack-go/slack"
)

// Slack support only a small set of markdown syntax.
// See: https://api.slack.com/reference/surfaces/formatting#basics
func NewMsgMarkdownFromPayload(payload *models.Payload) Msg {
	headerBlock := slack.NewHeaderBlock(
		slack.NewTextBlockObject(
			slack.PlainTextType,
			payload.Title,
			true, false,
		),
	)

	bodyBlock := slack.NewSectionBlock(
		slack.NewTextBlockObject(
			slack.MarkdownType,
			payload.Markdown,
			false, false,
		),
		nil,
		nil,
	)
	return []slack.Block{headerBlock, bodyBlock}
}
