package dingtalk

import "github.com/bougou/webhook-adapter/models"

const (
	MsgTypeLink = "link"
)

func init() {
	SupportedMsgtypes[MsgTypeLink] = NewMsgLinkFromPayload
}

type Link struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicURL     string `json:"picUrl"`
	MessageURL string `json:"messageUrl"`
}

func (*Link) dingTalkMsgtype() string {
	return MsgTypeLink
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
		MsgType: MsgTypeLink,
		Link:    link,
	}
}

func NewMsgLinkFromPayload(payload *models.Payload) *Msg {
	link := NewLink(payload.Title, payload.Text, "")
	return NewMsgLink(link)
}
