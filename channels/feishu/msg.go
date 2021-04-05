package feishu

import "time"

func init() {
	SupportedMsgtype["text"] = true      // 文本
	SupportedMsgtype["image"] = true     // 图片
	SupportedMsgtype["post"] = true      // 富文本（文章）
	SupportedMsgtype["sharechat"] = true // 群名片

	SupportedMsgtype["interactive"] = true // 消息卡片
}

type Msg struct {
	Timestamp time.Time `json:"timestamp"`
	Sign      string    `json:"sign"`

	MsgType string `json:"msg_type"`

	Content *Content `json:"content"`

	Card        *Card  `json:"card,omitempty"`
	RootID      string `json:"root_id,omitempty"`      // 需要回复的消息的open_message_id
	UpdateMulti bool   `json:"update_multi,omitempty"` // 控制卡片是否是共享卡片(所有用户共享同一张消息卡片），默认为 false
}

type Content struct {
	Text        string `json:"text,omitempty"`
	ImageKey    string `json:"image,omitempty"`
	Post        *Post  `json:"post,omitempty"`
	ShareChatID string `json:"share_chat_id,omitempty"`
}
