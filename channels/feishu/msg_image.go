package feishu

import "github.com/bougou/webhook-adapter/models"

func NewMsgImage(imageKey string) *Msg {
	return &Msg{
		MsgType: "text",
		Content: &Content{
			ImageKey: imageKey,
		},
	}
}

func (bot *FeishuGroupBot) SendImage(imageKey string) error {
	// 先上传图片
	// bot.UploadFile("")

	msg := NewMsgImage(imageKey)
	return bot.Send(msg)
}

func NewMsgImageFromPayload(payload *models.Payload) *Msg {
	return NewMsgImage(payload.Text)
}
