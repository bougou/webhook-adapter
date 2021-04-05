package feishu

import "github.com/bougou/webhook-adapter/models"

func NewMsgShareChat(shareChatID string) *Msg {
	return &Msg{
		MsgType: "share_chat",
		Content: &Content{
			ShareChatID: shareChatID,
		},
	}
}

func (bot *FeishuGroupBot) SendShareChat(shareChatID string) error {
	msg := NewMsgShareChat(shareChatID)
	return bot.Send(msg)
}

func NewMsgShareChatFromPayload(payload *models.Payload) *Msg {
	return &Msg{
		MsgType: "share_chat",
	}
}
