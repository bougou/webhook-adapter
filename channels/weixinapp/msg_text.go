package weixinapp

import "github.com/bougou/webhook-adapter/models"

const (
	MsgTypeText = "text"
)

func init() {
	SupportedMsgtypes[MsgTypeText] = NewMsgTextFromPayload
}

type Text struct {
	Content string `json:"content"`
}

func NewMsgText(content string) *Msg {
	return &Msg{
		MsgType: MsgTypeText,
		Text: &Text{
			Content: content,
		},
	}
}

func NewMsgTextFromPayload(payload *models.Payload) *Msg {
	return NewMsgText(payload.Text)
}
