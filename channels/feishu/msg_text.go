package feishu

import "github.com/bougou/webhook-adapter/models"

const (
	MsgTypeText = "text"
)

func init() {
	SupportedMsgtypes[MsgTypeText] = NewMsgTextFromPayload
}

func NewMsgText(text string) *Msg {
	return &Msg{
		MsgType: MsgTypeText,
		Content: &Content{
			Text: text,
		},
	}
}

func NewMsgTextFromPayload(payload *models.Payload) *Msg {
	return NewMsgText(payload.Text)
}
