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
	payload2Msg, ok := SupportedMsgtypes[s.msgType]
	if !ok {
		return fmt.Errorf("unkown msg type for slack")
	}
	msg := payload2Msg(payload)
	return s.bot.send(msg)
}
