package feishu

import (
	"github.com/bougou/webhook-adapter/channels/feishu/card"
	"github.com/bougou/webhook-adapter/models"
)

const (
	// Underlying, we use interactive msg type to implement markdown
	MsgTypeMarkdown = "markdown"
)

func init() {
	SupportedMsgtypes[MsgTypeMarkdown] = NewMsgMarkdownFromPayload
}

func NewMsgMarkdownFromPayload(payload *models.Payload) *Msg {
	card := NewCardMarkdown(payload.Title, payload.Markdown)

	return NewMsgCard(card)
}

func NewCardMarkdown(title string, markdown string) *Card {
	elements := []card.CardModule{}

	// see: https://open.feishu.cn/document/ukTMukTMukTM/uADOwUjLwgDM14CM4ATN

	// module := &card.ModuleDiv{
	// 	Tag: "div",
	// 	Text: &card.Text{
	// 		Tag:     "lark_md",
	// 		Content: markdown,
	// 	},
	// }

	module := card.NewModuleMarkdown(markdown, nil)

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
