package dingtalk

import "github.com/bougou/webhook-adapter/models"

type FeedCard struct {
	Links []*FeedCardLink `json:"links"`
}

func (*FeedCard) dingTalkMsgtype() string {
	return MsgTypeFeedCard
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
		MsgType:  MsgTypeFeedCard,
		FeedCard: feedCard,
	}
}

func NewMsgFeedCardFromPayload(payload *models.Payload) *Msg {
	// Todo, construct feedCardLinks from payload
	feedCardLinks := []*FeedCardLink{}
	feedCard := NewFeedCard(feedCardLinks)
	return NewMsgFeedCard(feedCard)
}
