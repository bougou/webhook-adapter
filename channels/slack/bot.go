package slack

import (
	"errors"
	"fmt"

	"github.com/slack-go/slack"
)

const ChannelTypeSlack = "slack"

var SupportedMsgtype = make(map[string]bool)

func ValidMsgtype(msgtype string) bool {
	if _, exists := SupportedMsgtype[msgtype]; !exists {
		return false
	}
	return true
}

// SlackBot can send messages to slack channel
type SlackBot struct {
	token   string
	channel string

	client *slack.Client
}

func NewSlackBot(token string, channel string) *SlackBot {
	client := slack.New(token, slack.OptionDebug(true))

	r, err := client.AuthTest()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("authresp:", r)

	return &SlackBot{
		token:   token,
		channel: channel,
		client:  client,
	}
}

func (s *SlackBot) send(msg Msg) error {
	_, _, err := s.client.PostMessage(
		s.channel,
		slack.MsgOptionBlocks(msg...),
	)
	if err != nil {
		msg := fmt.Sprintf("slack send failed, err: %s", err)
		return errors.New(msg)
	}
	return nil
}
