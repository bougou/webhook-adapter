package weixinapp

import "github.com/bougou/webhook-adapter/models"

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
