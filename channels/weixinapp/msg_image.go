package weixinapp

import "github.com/bougou/webhook-adapter/models"

const (
	MsgTypeImage = "image"
)

func init() {
	SupportedMsgtypes[MsgTypeImage] = NewMsgImageFromPayload
}

type Image struct {
	MediaID string `json:"media_id"`
}

func NewMsgImage(mediaID string) *Msg {

	return &Msg{
		MsgType: MsgTypeImage,
		Image: &Image{
			MediaID: mediaID,
		},
	}
}

func NewMsgImageFromPayload(payload *models.Payload) *Msg {
	// Todo

	mediaID := ""
	return NewMsgImage(mediaID)
}
