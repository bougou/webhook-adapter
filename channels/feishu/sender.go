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
	return &Sender{
		bot:     NewFeishuGroupBot(token),
		msgType: msgType,
	}
}

func (s *Sender) Send(payload *models.Payload) error {
	var msg *Msg

	switch s.msgType {
	case "text":
		msg = NewMsgTextFromPayload(payload)
	case "image":
		msg = NewMsgImageFromPayload(payload)
	case "post":
		msg = NewMsgPostFromPayload(payload)
	case "sharechat":
		msg = NewMsgShareChatFromPayload(payload)
	case "interactive":
		msg = NewMsgInteractiveFromPayload(payload)
	default:
		return fmt.Errorf("unkown msg type")
	}

	return s.bot.Send(msg)
}
