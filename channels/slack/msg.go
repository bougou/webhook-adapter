package slack

import (
	"github.com/slack-go/slack"
)

func init() {
	SupportedMsgtype[MsgTypeText] = true
	SupportedMsgtype[MsgTypeMarkdown] = true
}

const (
	MsgTypeText     = "text"
	MsgTypeMarkdown = "markdown"
)

type Msg []slack.Block
