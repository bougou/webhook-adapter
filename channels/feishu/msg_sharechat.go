package feishu

import "github.com/bougou/webhook-adapter/models"

const (
	MsgTypeShareChat = "sharechat"
)

func init() {
	SupportedMsgtypes[MsgTypeShareChat] = NewMsgShareChatFromPayload
}

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
