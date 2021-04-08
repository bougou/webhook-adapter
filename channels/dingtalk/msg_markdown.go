package dingtalk

import "github.com/bougou/webhook-adapter/models"

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func (*Markdown) dingTalkMsgtype() string {
	return MsgTypeMarkdown
}

func (md *Markdown) Valid() bool {
	if md.Title == "" || md.Text == "" {
		return false
	}
	return true
}

func NewMarkdown(title string, text string) *Markdown {
	return &Markdown{title, text}
}

func NewMsgMarkdown(md *Markdown) *Msg {
	return &Msg{
		MsgType:  MsgTypeMarkdown,
		Markdown: md,
	}
}

func NewMsgMarkdownFromPayload(payload *models.Payload) *Msg {
	md := NewMarkdown(payload.Title, payload.Markdown)
	msg := NewMsgMarkdown(md)
	msg.WithAtMobiles(payload.At.AtMobiles)
	msg.WithAtAll(payload.At.AtAll)

	return msg
}
