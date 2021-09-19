package dingtalk

import "github.com/bougou/webhook-adapter/models"

const (
	MsgTypeActionCard = "actioncard"
)

func init() {
	SupportedMsgtypes[MsgTypeActionCard] = NewMsgActionCardFromPayload
}

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
	return MsgTypeActionCard
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

func NewMsgActionCard(actionCard *ActionCard) *Msg {
	return &Msg{
		MsgType:    MsgTypeActionCard,
		ActionCard: actionCard,
	}
}

func NewMsgActionCardFromPayload(p *models.Payload) *Msg {
	// Todo, construct actionCard from payload
	actionCard := &ActionCard{}
	return &Msg{
		MsgType:    MsgTypeActionCard,
		ActionCard: actionCard,
	}
}
