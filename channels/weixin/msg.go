package weixin

const (
	MsgTypeFile     = "file"
	MsgTypeImage    = "image"
	MsgTypeMarkdown = "markdown"
	MsgTypeNews     = "news"
	MsgTypeText     = "text"
)

func init() {
	SupportedMsgtype[MsgTypeFile] = true
	SupportedMsgtype[MsgTypeImage] = true
	SupportedMsgtype[MsgTypeMarkdown] = true
	SupportedMsgtype[MsgTypeNews] = true
	SupportedMsgtype[MsgTypeText] = true

}

type Msg struct {
	MsgType  string    `json:"msgtype"`
	File     *File     `json:"file,omitempty"`
	Image    *Image    `json:"image,omitempty"`
	Markdown *Markdown `json:"markdown,omitempty"`
	News     *News     `json:"news,omitempty"`
	Text     *Text     `json:"text,omitempty"`
}
