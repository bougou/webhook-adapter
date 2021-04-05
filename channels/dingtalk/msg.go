package dingtalk

func init() {
	SupportedMsgtype["text"] = true
	SupportedMsgtype["markdown"] = true
	SupportedMsgtype["link"] = true
	SupportedMsgtype["feedcard"] = true
	SupportedMsgtype["actioncard"] = true
}

type Msg struct {
	MsgType    string      `json:"msgtype"`
	Text       *Text       `json:"text,omitempty"`
	Link       *Link       `json:"link,omitempty"`
	Markdown   *Markdown   `json:"markdown,omitempty"`
	ActionCard *ActionCard `json:"actionCard,omitempty"`
	FeedCard   *FeedCard   `json:"feedCard,omitempty"`
	At         *At         `json:"at,omitempty"` // only available for text and markdown type
}

type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

func (msg *Msg) SupportAt() bool {
	if msg.MsgType == "text" || msg.MsgType == "markdown" {
		return true
	}
	return false
}

func (msg *Msg) WithAt(at *At) *Msg {
	if msg.SupportAt() {
		msg.At = at
	}
	return msg
}

func (msg *Msg) WithAtAll(atAll bool) *Msg {
	if msg.SupportAt() && msg.At == nil {
		msg.At = &At{}
	}
	msg.At.IsAtAll = atAll
	return msg
}

func (msg *Msg) WithAtMobiles(mobiles []string) *Msg {
	if msg.SupportAt() && msg.At == nil {
		msg.At = &At{}
	}
	msg.At.AtMobiles = mobiles
	return msg
}
