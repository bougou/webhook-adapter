package feishu

const (
	MsgTypeText        = "text"
	MsgTypeImage       = "image"
	MsgTypePost        = "post"
	MsgTypeShareChat   = "sharechat"
	MsgTypeInteractive = "interactive"

	// Underlying, we use interactive to implement markdown
	MsgTypeMarkdown = "markdown"
)

func init() {
	SupportedMsgtype[MsgTypeText] = true        // 文本
	SupportedMsgtype[MsgTypeImage] = true       // 图片
	SupportedMsgtype[MsgTypePost] = true        // 富文本（文章）
	SupportedMsgtype[MsgTypeShareChat] = true   // 群名片
	SupportedMsgtype[MsgTypeInteractive] = true // 消息卡片
	SupportedMsgtype[MsgTypeMarkdown] = true    // markdown 文本
}

type Msg struct {
	// 开启签名验证后发送文本消息
	// Timestamp time.Time `json:"timestamp,omitempty"`
	// Sign      string    `json:"sign,omitempty"`

	MsgType string `json:"msg_type"`

	Content *Content `json:"content,omitempty"`

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
