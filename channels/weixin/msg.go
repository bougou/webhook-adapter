package weixin

import "github.com/bougou/webhook-adapter/models"

const (
	ChannelTypeWeixin string = "weixin"

	MsgTypeFile string = "file"

	MsgTypeImage string = "image"
	maxImageSize int64  = 2 * 1024 * 1024 // 2MB

	MsgTypeMarkdown  string = "markdown"
	maxMarkdownBytes int    = 4096
	truncatedMark    string = "\n... more is truncated due to limit"

	MsgTypeNews         string = "news"
	maxArticlesNumber   int    = 8
	maxTitleBytes       int    = 128
	maxDescriptionBytes int    = 512

	MsgTypeTemplateCard string = "template_card"

	MsgTypeText  string = "text"
	maxTextBytes int    = 2048
)

type Msg struct {
	MsgType      string        `json:"msgtype"`
	File         *File         `json:"file,omitempty"`
	Image        *Image        `json:"image,omitempty"`
	Markdown     *Markdown     `json:"markdown,omitempty"`
	News         *News         `json:"news,omitempty"`
	Text         *Text         `json:"text,omitempty"`
	TemplateCard *TemplateCard `json:"template_card,omitempty"`
}

type Payload2MsgFn func(payload *models.Payload) *Msg

var Payload2MsgFnMap = make(map[string]Payload2MsgFn)

func ValidMsgtype(msgtype string) bool {
	if _, exists := Payload2MsgFnMap[msgtype]; !exists {
		return false
	}

	return true
}

func ValidMsg(msgType string, msg *Msg) error {
	return nil
}
