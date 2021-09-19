package weixinapp

import "github.com/bougou/webhook-adapter/models"

const (
	MsgTypeFile = "file"
)

func init() {
	SupportedMsgtypes[MsgTypeFile] = NewMsgFileFromPayload
}

type File struct {
	MediaID string `json:"media_id"`
}

func NewMsgFile(file *File) *Msg {
	return &Msg{
		MsgType: MsgTypeFile,
		File:    file,
	}
}

func NewMsgFileFromPayload(payload *models.Payload) *Msg {
	msg := &Msg{}

	return msg
}
