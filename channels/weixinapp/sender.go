package weixinapp

import (
	"fmt"

	"github.com/bougou/webhook-adapter/models"
)

type Sender struct {
	notifer *Notifier
	msgType string
}

func NewSender(corpID string, agentID int, agentSecret string, msgType string) *Sender {
	return &Sender{
		notifer: NewNotifer(corpID, agentID, agentSecret),
		msgType: msgType,
	}
}

func (s *Sender) Send(payload *models.Payload) error {
	var msg *Msg

	switch s.msgType {
	case "text":
		msg = NewMsgTextFromPayload(payload)
	case "file":
		msg = NewMsgFileFromPayload(payload)
	case "markdown":
		msg = NewMsgMarkdownFromPayload(payload)
	case "image":
		msg = NewMsgImageFromPayload(payload)
	case "news":
		msg = NewMsgNewsFromPayload(payload)
	default:
		return fmt.Errorf("unkown msg type")
	}

	return s.notifer.Send(msg)
}
