package weixinapp

import (
	"fmt"

	"github.com/bougou/webhook-adapter/models"
)

type Sender struct {
	notifer *Notifier
	msgType string
}

func NewSender(corpID string, agentID int, agentSecret string, msgType string, toUser string, toParty string, toTag string) *Sender {
	return &Sender{
		notifer: NewNotifer(corpID, agentID, agentSecret, toUser, toParty, toTag),
		msgType: msgType,
	}
}

func (s *Sender) Send(payload *models.Payload) error {
	payload2Msg, ok := SupportedMsgtypes[s.msgType]
	if !ok {
		return fmt.Errorf("unkown msg type")
	}
	msg := payload2Msg(payload)
	return s.notifer.Send(msg)
}
