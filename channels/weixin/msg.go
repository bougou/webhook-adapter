package weixin

type Msg struct {
	MsgType  string    `json:"msgtype"`
	File     *File     `json:"file,omitempty"`
	Image    *Image    `json:"image,omitempty"`
	Markdown *Markdown `json:"markdown,omitempty"`
	News     *News     `json:"news,omitempty"`
	Text     *Text     `json:"text,omitempty"`
}
