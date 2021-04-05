package weixin

import (
	"fmt"
	"io"

	"github.com/bougou/webhook-adapter/models"
)

type File struct {
	MediaID string `json:"media_id"`
}

func NewMsgFile(mediaID string) *Msg {
	return &Msg{
		MsgType: "file",
		File: &File{
			MediaID: mediaID,
		},
	}
}

func (b *WeixinGroupBot) SendFile(filename string, fileReader io.Reader) error {
	mediaID, err := b.UploadFile(filename, fileReader)
	if err != nil {
		return fmt.Errorf("send file error, err: %v", err)
	}

	msg := NewMsgFile(mediaID)
	return b.Send(msg)
}

func NewMsgFileFromPayload(payload *models.Payload) *Msg {
	return &Msg{
		MsgType: "file",
	}
}
