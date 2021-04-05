package dingtalk

import "github.com/bougou/webhook-adapter/models"

type Text struct {
	Content string `json:"content"`
}

func (*Text) dingTalkMsgtype() string {
	return "text"
}

func NewText(content string) *Text {
	return &Text{content}
}

func NewMsgText(text *Text) *Msg {
	return &Msg{
		MsgType: "text",
		Text:    text,
	}
}

func (bot *DingtalkGroupBot) SendText(content string, atMobiles []string, atAll bool) error {
	text := NewText(content)
	msg := NewMsgText(text)
	msg.WithAtAll(atAll).WithAtMobiles(atMobiles)
	return bot.send(msg)
}

func NewMsgTextFromPayload(payload *models.Payload) *Msg {
	return &Msg{
		MsgType: "text",
		Text:    NewText(payload.Text),
	}
}
