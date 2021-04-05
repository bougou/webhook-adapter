package feishu

import "github.com/bougou/webhook-adapter/models"

func NewMsgText(text string) *Msg {
	return &Msg{
		MsgType: "text",
		Content: &Content{
			Text: text,
		},
	}
}

func (bot *FeishuGroupBot) SendText(text string) error {
	msg := NewMsgText(text)
	return bot.Send(msg)
}

func NewMsgTextFromPayload(payload *models.Payload) *Msg {
	return NewMsgText(payload.Text)
}
