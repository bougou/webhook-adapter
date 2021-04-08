package feishu

import (
	"github.com/bougou/webhook-adapter/channels/feishu/card"
	"github.com/bougou/webhook-adapter/models"
)

type Card struct {
	Config       *CardConfig       `json:"config"`
	Header       *CardHeader       `json:"header"`
	CardLink     *card.MultiURL    `json:"card_link"`
	Elements     []card.CardModule `json:"elements"` // 最多可堆叠 50 个模块
	I18NElements *I18NElements     `json:"i18n_elements"`
}

// CardConfig 卡片配置
type CardConfig struct {
	WideScreenMode bool `json:"wide_screen_mode"` // 2021/03/22 之后，此字段废弃，所有卡片均升级为自适应屏幕宽度的宽版卡片
	EnableForward  bool `json:"enable_forward"`   // 是否允许卡片被转发，默认 false
}

type CardHeader struct {
	Title    *card.Text `json:"title"`              // 卡片标题内容, text 对象（仅支持 "plain_text")
	Template string     `json:"template,omitempty"` // 控制标题背景颜色, https://open.feishu.cn/document/ukTMukTMukTM/ukTNwUjL5UDM14SO1ATN
}

type I18NElements struct {
	ZHCN []card.CardModule `json:"zh_cn"`
	ENUS []card.CardModule `json:"en_us"`
	JAJP []card.CardModule `json:"jn_jp"`
}

func NewMsgCard(card *Card) *Msg {
	return &Msg{
		MsgType: MsgTypeInteractive,
		Card:    card,
	}
}

func NewMsgInteractiveFromPayload(payload *models.Payload) *Msg {
	// Todo, construct card from payload
	card := &Card{}

	return NewMsgCard(card)
}

func NewMsgMarkdownFromPayload(payload *models.Payload) *Msg {
	card := NewCardMarkdown(payload.Title, payload.Markdown)

	return NewMsgCard(card)
}

func NewCardMarkdown(title string, markdown string) *Card {
	elements := []card.CardModule{}

	module := &card.ModuleDiv{
		Tag: "div",
		Text: &card.Text{
			Tag:     "lark_md",
			Content: markdown,
		},
	}
	elements = append(elements, module)

	return &Card{
		Config: &CardConfig{
			EnableForward: false,
		},
		Header: &CardHeader{
			Title: &card.Text{
				Tag:     "plain_text",
				Content: title,
			},
		},
		Elements: elements,
	}
}
