package slack

import (
	"fmt"

	"github.com/bougou/webhook-adapter/models"
)

type Sender struct {
	bot     *SlackBot
	msgType string
}

func NewSender(token string, channel string, msgType string) *Sender {
	if msgType == "" {
		msgType = MsgTypeMarkdown
	}

	return &Sender{
		bot:     NewSlackBot(token, channel),
		msgType: msgType,
	}
}

func (s *Sender) Send(payload *models.Payload) error {
	var msg Msg
	switch s.msgType {
	case "text":
		msg = NewMsgTextFromPayload(payload)
	case "markdown":
		msg = NewMsgMarkdownFromPayload(payload)
	default:
		return fmt.Errorf("unkown msg type")
	}

	return s.bot.send(msg)
}
