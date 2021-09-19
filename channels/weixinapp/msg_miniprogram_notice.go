package weixinapp

import "github.com/bougou/webhook-adapter/models"

const (
	MsgTypeMiniprogramNotice = "miniprogram_notice"
)

func init() {
	SupportedMsgtypes[MsgTypeMiniprogramNotice] = NewMsgMiniprogramNoticeFromPayload
}

type MiniprogramNotice struct {
	AppID             string `json:"appid"`
	Page              string `json:"page"`
	Title             string `json:"title"`
	Description       string `json:"description"`
	EmphasisFirstItem string `json:"emphasis_first_item"`
	ContentItem       []KV   `json:"content_item"`
}

type KV struct {
	Key   string `json:"key"`
	Value string `json:"string"`
}

// Todo
func NewMsgMiniprogramNoticeFromPayload(payload *models.Payload) *Msg {
	return &Msg{}
}
