package weixinapp

import (
	"fmt"
	"strings"

	"github.com/bougou/webhook-adapter/models"
	"github.com/bougou/webhook-adapter/utils"
)

const (
	maxMarkdownBytes int = 2048
	truncatedMark        = "\n... more is truncated due to limit"
	// see: https://work.weixin.qq.com/api/doc/90000/90135/90236#markdown%E6%B6%88%E6%81%AF
	MsgTypeMarkdown = "markdown"
)

func init() {
	SupportedMsgtypes[MsgTypeMarkdown] = NewMsgMarkdownFromPayload
}

type Markdown struct {
	Content string `json:"content"` // this should be raw markdown string, weixin bot only support a small subset syntax
}

func NewMsgMarkdown(content string) *Msg {
	content = SanitizeMarkdown(content)
	truncated := utils.TruncateToValidUTF8(content, maxMarkdownBytes, truncatedMark)

	return &Msg{
		MsgType: MsgTypeMarkdown,
		Markdown: &Markdown{
			Content: truncated,
		},
	}
}

func NewMsgMarkdownFromPayload(payload *models.Payload) *Msg {
	m := fmt.Sprintf("%s\n%s", payload.Title, payload.Markdown)
	return NewMsgMarkdown(m)
}

func SanitizeMarkdown(content string) string {
	// no need <br> for line break
	content = strings.ReplaceAll(content, "<br>", "")

	// remove `>` line
	content = strings.ReplaceAll(content, "\n>\n", "\n")

	return content
}
