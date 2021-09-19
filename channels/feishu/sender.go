package feishu

import (
	"fmt"

	"github.com/bougou/webhook-adapter/models"
)

type Sender struct {
	bot     *FeishuGroupBot
	msgType string
}

func NewSender(token string, msgType string) *Sender {
	if msgType == "" {
		msgType = MsgTypeMarkdown
	}

	return &Sender{
		bot:     NewFeishuGroupBot(token),
		msgType: msgType,
	}
}

func (s *Sender) Send(payload *models.Payload) error {
	payload2Msg, ok := SupportedMsgtypes[s.msgType]
	if !ok {
		return fmt.Errorf("unkown msg type")
	}
	msg := payload2Msg(payload)
	return s.bot.Send(msg)
}
