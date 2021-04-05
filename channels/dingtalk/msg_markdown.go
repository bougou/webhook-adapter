package dingtalk

import "github.com/bougou/webhook-adapter/models"

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func (*Markdown) dingTalkMsgtype() string {
	return "markdown"
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
		MsgType:  "markdown",
		Markdown: md,
	}
}

func (bot *DingtalkGroupBot) SendMarkdown(title string, text string, atMobiles []string, atAll bool) error {
	md := NewMarkdown(title, text)
	msg := NewMsgMarkdown(md)
	msg.WithAtAll(atAll).WithAtMobiles(atMobiles)
	return bot.send(msg)
}

func NewMsgMarkdownFromPayload(payload *models.Payload) *Msg {
	return &Msg{
		MsgType: "feedCard",
	}
}
