package dingtalk

import "github.com/bougou/webhook-adapter/models"

type ActionCard struct {
	Title          string `json:"title"`          // seems no meaning now
	Text           string `json:"text"`           // support markdown format
	Btnorientation string `json:"btnOrientation"` // 0：按钮竖直排列, 1：按钮横向排列
	Hideavatar     string `json:"hideAvatar,omitempty"`

	Singletitle string `json:"singleTitle,omitempty"`
	Singleurl   string `json:"singleURL,omitempty"`

	Btns []*Btn `json:"btns,omitempty"`
}

func (*ActionCard) dingTalkMsgtype() string {
	return "actionCard"
}

type Btn struct {
	Title     string `json:"title"`
	Actionurl string `json:"actionURL"`
}

func NewBtn(title, url string) *Btn {
	return &Btn{title, url}
}

func NewActionCard(title string, text string, horizonBtn bool, btns []*Btn) *ActionCard {
	actionCard := &ActionCard{
		Title: title,
		Text:  text,
	}

	if horizonBtn {
		actionCard.Btnorientation = "1" // 按钮横向排列
	} else {
		actionCard.Btnorientation = "0" // 按钮竖直排列
	}

	if len(btns) == 1 {
		actionCard.Singletitle = btns[0].Title
		actionCard.Singletitle = btns[0].Actionurl
	} else {
		actionCard.Btns = btns
	}

	return actionCard
}

func NewMsgActionCardFromPayload(p *models.Payload) *Msg {

	actionCard := &ActionCard{}
	return &Msg{
		MsgType:    "actionCard",
		ActionCard: actionCard,
	}

}

func NewMsgActionCard(actionCard *ActionCard) *Msg {
	return &Msg{
		MsgType:    "actionCard",
		ActionCard: actionCard,
	}
}

func (bot *DingtalkGroupBot) SendActionCard(title string, text string, horizonBtn bool, btns []*Btn) error {
	actionCard := NewActionCard(title, text, horizonBtn, btns)
	msg := NewMsgActionCard(actionCard)
	return bot.send(msg)
}
