package weixinapp

import "github.com/bougou/webhook-adapter/models"

const (
	MsgTypeVoice = "voice"
)

func init() {
	SupportedMsgtypes[MsgTypeVoice] = NewMsgVoiceFromPayload
}

type Voice struct {
	MediaID string `json:"media_id"`
}

// Todo
func NewMsgVoiceFromPayload(payload *models.Payload) *Msg {
	return &Msg{}
}
