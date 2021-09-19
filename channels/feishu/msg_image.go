package feishu

import "github.com/bougou/webhook-adapter/models"

const (
	MsgTypeImage = "image"
)

func init() {
	SupportedMsgtypes[MsgTypeImage] = NewMsgImageFromPayload
}

func NewMsgImage(imageKey string) *Msg {
	return &Msg{
		MsgType: MsgTypeImage,
		Content: &Content{
			ImageKey: imageKey,
		},
	}
}

func (bot *FeishuGroupBot) SendImage(imageKey string) error {

	msg := NewMsgImage(imageKey)
	return bot.Send(msg)
}

func NewMsgImageFromPayload(payload *models.Payload) *Msg {

	// Todo 先上传图片得到 imageKey
	// bot.UploadFile("")

	imageKey := ""
	return NewMsgImage(imageKey)
}
