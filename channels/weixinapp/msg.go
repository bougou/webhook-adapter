package weixinapp

const (
	MsgTypeText              = "text"
	MsgTypeFile              = "file"
	MsgTypeImage             = "image"
	MsgTypeVoice             = "voice"
	MsgTypeVideo             = "video"
	MsgTypeNews              = "news"
	MsgTypeMPNews            = "mpnews"
	MsgTypeTextCard          = "textcard"
	MsgTypeMarkdown          = "markdown"
	MsgTypeMiniprogramNotice = "miniprogram_notice"
	MsgTypeInteractiveCard   = "interactive_card"
)

func init() {
	SupportedMsgtype[MsgTypeText] = true
	SupportedMsgtype[MsgTypeFile] = true
	SupportedMsgtype[MsgTypeImage] = true
	SupportedMsgtype[MsgTypeVoice] = true
	SupportedMsgtype[MsgTypeVideo] = true
	SupportedMsgtype[MsgTypeNews] = true
	SupportedMsgtype[MsgTypeMPNews] = true
	SupportedMsgtype[MsgTypeTextCard] = true
	SupportedMsgtype[MsgTypeMarkdown] = true
	SupportedMsgtype[MsgTypeMiniprogramNotice] = true
	SupportedMsgtype[MsgTypeInteractiveCard] = true
}

type Msg struct {
	MsgType string `json:"msgtype"`
	AgentID int    `json:"agentid"`

	// touser、toparty、totag不能同时为空
	ToUser  string `json:"touser,omitempty"`
	ToParty string `json:"toparty,omitempty"`
	ToTag   string `json:"totag,omitempty"`

	Text                *Text                `json:"text,omitempty"`
	Image               *Image               `json:"image,omitempty"`
	Voice               *Voice               `json:"voice,omitempty"`
	File                *File                `json:"file,omitempty"`
	TextCard            *TextCard            `json:"textcard,omitempty"`
	News                *News                `json:"news,omitempty"`
	MPNews              *MPNews              `json:"mpnews,omitempty"`
	Markdown            *Markdown            `json:"markdown,omitempty"`
	MiniprogramNotice   *MiniprogramNotice   `json:"miniprogram_notice,omitempty"`
	InteractiveTaskcard *InteractiveTaskcard `json:"interactive_taskcard,omitempty"`

	Safe                   int `json:"safe,omitempty"`
	EnableIDTrans          int `json:"enable_id_trans,omitempty"`
	DuplicateCheckInterval int `json:"duplicate_check_interval,omitempty"`
}

func (msg *Msg) Valid() bool {
	if msg.ToUser == "" && msg.ToParty == "" && msg.ToTag == "" {
		return false
	}

	return true
}
