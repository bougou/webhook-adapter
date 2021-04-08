package weixin

import (
	"fmt"

	"github.com/bougou/webhook-adapter/models"
)

type Sender struct {
	bot     *WeixinGroupBot
	msgType string
}

func NewSender(key string, msgType string) *Sender {
	if msgType == "" {
		msgType = MsgTypeMarkdown
	}

	return &Sender{
		bot:     NewWexinGroupBot(key),
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

	return s.bot.Send(msg)

}
