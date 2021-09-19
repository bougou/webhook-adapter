package channels

import (
	"github.com/bougou/webhook-adapter/channels/dingtalk"
	"github.com/bougou/webhook-adapter/channels/feishu"
	"github.com/bougou/webhook-adapter/channels/slack"
	"github.com/bougou/webhook-adapter/channels/weixin"
	"github.com/bougou/webhook-adapter/channels/weixinapp"
	"github.com/bougou/webhook-adapter/models"
)

func NewDingtalkSender(token string, msgType string) models.Sender {
	return dingtalk.NewSender(token, msgType)
}

func NewFeishuSender(token string, msgType string) models.Sender {
	return feishu.NewSender(token, msgType)
}

func NewWeixinSender(token string, msgType string) models.Sender {
	return weixin.NewSender(token, msgType)
}

func NewWeixinAppSender(corpID string, agentID int, agentSecret string, msgType string) models.Sender {
	return weixinapp.NewSender(corpID, agentID, agentSecret, msgType)
}

func NewSlackSender(token string, channel string, msgType string) models.Sender {
	return slack.NewSender(token, channel, msgType)
}
