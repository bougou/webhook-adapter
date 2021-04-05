package dingtalk

import "github.com/bougou/webhook-adapter/models"

type FeedCard struct {
	Links []*FeedCardLink `json:"links"`
}

func (*FeedCard) dingTalkMsgtype() string {
	return "feedcard"
}

type FeedCardLink struct {
	Title      string `json:"title"`
	MessageURL string `json:"messageURL"`
	PicURL     string `json:"picURL"`
}

func NewFeedCardLink(title, messageURl, picURL string) *FeedCardLink {
	return &FeedCardLink{title, messageURl, picURL}
}

func NewFeedCard(links []*FeedCardLink) *FeedCard {
	return &FeedCard{
		Links: links,
	}
}

func NewMsgFeedCard(feedCard *FeedCard) *Msg {
	return &Msg{
		MsgType:  "feedCard",
		FeedCard: feedCard,
	}
}

func (bot *DingtalkGroupBot) SendFeedCard(links []*FeedCardLink) error {
	feedCard := NewFeedCard(links)
	msg := NewMsgFeedCard(feedCard)
	return bot.send(msg)
}

func NewMsgFeedCardFromPayload(payload *models.Payload) *Msg {
	return &Msg{
		MsgType: "feedCard",
	}
}
