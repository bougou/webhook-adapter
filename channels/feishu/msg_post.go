package feishu

import "github.com/bougou/webhook-adapter/models"

type Post struct {
	*ZHCN `json:"zh_ch"`
}

type ZHCN struct {
	Title   string  `json:"title"`   // 文章标题
	Content []*Line `json:"content"` // 文章内容，有多个行组成
}

type Line []*Segment // 每行由多个片段组成

type Segment struct {
	Tag      string `json:"tag"` // text, img, a, at
	UnEscape bool   `json:"un_escape"`
	Text     string `json:"text"`
	Href     string `json:"href"`
	UserID   string `json:"user_id"`
	ImageKey string `json:"image_key"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

func NewSegmentText(text string, unescape bool) *Segment {
	return &Segment{
		Tag:      "text",
		Text:     text,
		UnEscape: unescape,
	}
}

func NewSegmentA(text string, unescape bool, href string) *Segment {
	return &Segment{
		Tag:      "a",
		Text:     text,
		UnEscape: unescape,
		Href:     href,
	}
}

func NewSegmentAt(userID string) *Segment {
	return &Segment{
		Tag:    "at",
		UserID: userID,
	}
}

func NewSegmentImg(imageKey string, height int, width int) *Segment {
	return &Segment{
		Tag:      "img",
		ImageKey: imageKey,
		Height:   height,
		Width:    width,
	}
}

func NewMsgPost(title string, lines []*Line) *Msg {
	post := &Post{
		ZHCN: &ZHCN{
			Title:   title,
			Content: lines,
		},
	}

	return &Msg{
		MsgType: "post",
		Content: &Content{
			Post: post,
		},
	}
}
func (bot *FeishuGroupBot) SendPost(title string, lines []*Line) error {
	// 先上传图片
	// bot.UploadFile("")

	msg := NewMsgPost(title, lines)
	return bot.Send(msg)
}

func NewMsgPostFromPayload(payload *models.Payload) *Msg {
	return &Msg{
		MsgType: "post",
	}
}
