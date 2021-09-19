package weixinapp

import "github.com/bougou/webhook-adapter/models"

const (
	MsgTypeTextCard = "textcard"
)

func init() {
	SupportedMsgtypes[MsgTypeTextCard] = NewMsgTextCardFromPayload
}

type TextCard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	BtnText     string `json:"btntext"`
}

// Todo
func NewMsgTextCardFromPayload(payload *models.Payload) *Msg {
	return &Msg{}
}
