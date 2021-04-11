package weixinapp

import "github.com/bougou/webhook-adapter/models"

const maxMarkdownBytes int = 2048
const truncatedMark = "\n... more is truncated due to limit"

type Markdown struct {
	Content string `json:"content"` // this should be raw markdown string, weixin bot only support a small subset syntax
}

func NewMsgMarkdown(content string) *Msg {
	return &Msg{
		MsgType: MsgTypeMarkdown,
		Markdown: &Markdown{
			Content: content,
		},
	}
}

func NewMsgMarkdownFromPayload(payload *models.Payload) *Msg {

	return NewMsgMarkdown(payload.Markdown)
}
