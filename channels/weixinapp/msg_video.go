package weixinapp

import "github.com/bougou/webhook-adapter/models"

const (
	MsgTypeVideo = "video"
)

func init() {
	SupportedMsgtypes[MsgTypeVideo] = NewMsgVideoFromPayload
}

type Video struct {
	MediaID     string `json:"media_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Todo
func NewMsgVideoFromPayload(payload *models.Payload) *Msg {
	return &Msg{}
}
