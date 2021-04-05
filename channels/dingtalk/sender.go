package dingtalk

import (
	"fmt"

	"github.com/bougou/webhook-adapter/models"
)

type Sender struct {
	bot     *DingtalkGroupBot
	msgType string
}

func NewSender(token string, msgType string) *Sender {
	return &Sender{
		bot:     NewDingtalkGroupBot(token),
		msgType: msgType,
	}
}

func (s *Sender) Send(payload *models.Payload) error {
	fmt.Println("dingtak sender")
	var msg *Msg

	switch s.msgType {
	case "text":
		msg = NewMsgTextFromPayload(payload)
	case "link":
		msg = NewMsgLinkFromPayload(payload)
	case "markdown":
		msg = NewMsgMarkdownFromPayload(payload)
	case "feedcard":
		msg = NewMsgFeedCardFromPayload(payload)
	case "actioncard":
		msg = NewMsgActionCardFromPayload(payload)
	default:
		return fmt.Errorf("unkown msg type")
	}

	return s.bot.send(msg)
}
