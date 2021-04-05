package dingtalk

import "github.com/bougou/webhook-adapter/models"

type Link struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicURL     string `json:"picUrl"`
	MessageURL string `json:"messageUrl"`
}

func (*Link) dingTalkMsgtype() string {
	return "link"
}

func NewLink(title string, text string, messageURL string) *Link {
	return &Link{
		Text:       text,
		Title:      title,
		MessageURL: messageURL,
	}
}

func (link *Link) WithPicURL(picURL string) *Link {
	link.PicURL = picURL
	return link
}

func NewMsgLink(link *Link) *Msg {
	return &Msg{
		MsgType: "link",
		Link:    link,
	}
}

func (bot *DingtalkGroupBot) SendLink(title string, text string, messageURL string, picURL string) error {
	link := NewLink(title, text, messageURL)
	link.WithPicURL(picURL)
	msg := NewMsgLink(link)
	return bot.send(msg)
}

func NewMsgLinkFromPayload(payload *models.Payload) *Msg {
	return &Msg{
		MsgType: "feedCard",
	}
}
