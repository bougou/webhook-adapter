package weixin

import (
	"github.com/bougou/webhook-adapter/models"
)

const maxMarkdownBytes int = 4096
const truncatedMark = "\n... more is truncated due to limit"

type Markdown struct {
	Content string `json:"content"` // this should be raw markdown string, weixin bot only support a small subset syntax
}

// underscore(_) format shows different on phone app and desktop app
// suggest to use star(*) format
var rawSupportedMarkdown = `# header 1
## header 2
### header 3
#### header 4
##### header 5
[link](http.www.baidu.com)
<font color="info">info color</font>
<font color="comment">comment color</font>
<font color="warning">warning color</font>
plain text
***three star***
**two star**
*one star*
___three underscore___
__two underscore__
_one underscore_
> reference
> reference
`

func NewMsgMarkdown(content string) *Msg {
	truncated := TruncateToValidUTF8(content, maxMarkdownBytes, truncatedMark)
	msg := &Msg{
		MsgType: "markdown",
		Markdown: &Markdown{
			Content: truncated,
		},
	}
	return msg
}

func (b *WeixinGroupBot) SendMarkdown(content string) error {
	msg := NewMsgMarkdown(content)
	return b.Send(msg)
}

func NewMsgMarkdownFromPayload(payload *models.Payload) *Msg {
	return NewMsgMarkdown(payload.Markdown)
}
