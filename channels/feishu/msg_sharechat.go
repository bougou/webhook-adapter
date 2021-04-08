package feishu

import "github.com/bougou/webhook-adapter/models"

func NewMsgShareChat(shareChatID string) *Msg {
	return &Msg{
		MsgType: MsgTypeShareChat,
		Content: &Content{
			ShareChatID: shareChatID,
		},
	}
}

func NewMsgShareChatFromPayload(payload *models.Payload) *Msg {
	// todo

	shareChatID := ""
	return NewMsgShareChat(shareChatID)
}
