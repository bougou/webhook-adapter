package feishu

import "github.com/bougou/webhook-adapter/models"

type Card struct {
	Config        *CardConfig  `json:"config"`
	CardLink      *MultiURL    `json:"card_link"`
	Header        *CardHeader  `json:"header"`
	Elements      []CardModule `json:"elements"` // 最多可堆叠 50 个模块
	*I18NElements `json:"i18n_elements"`
}

// CardConfig 卡片配置
type CardConfig struct {
	WideScreenMode bool `json:"wide_screen_mode"` // 2021/03/22 之后，此字段废弃，所有卡片均升级为自适应屏幕宽度的宽版卡片
	EnableForward  bool `json:"enable_forward"`   // 是否允许卡片被转发，默认 false
}

type CardHeader struct {
	Title    Text   `json:"title"`              // 卡片标题内容, text 对象（仅支持 "plain_text")
	Template string `json:"template,omitempty"` // 控制标题背景颜色, https://open.feishu.cn/document/ukTMukTMukTM/ukTNwUjL5UDM14SO1ATN
}

type I18NElements struct {
	ZHCN []CardModule `json:"zh_cn"`
	ENUS []CardModule `json:"en_us"`
	JAJP []CardModule `json:"jn_jp"`
}

// div, hr, img, note, action
type CardModule interface {
	cardModule() string
}

// ModuleDiv 内容模块
type ModuleDiv struct {
	Tag    string   `json:"tag"`   // div
	Text   *Text    `json:"text"`  // 单个文本展示, 和 field 至少要有一个
	Fields []Field  `json:"field"` // 多个文本展示, 和 text 至少要有一个
	Extra  *Element `json:"extra"` // 展示附加元素, 最多可展示一个元素
}

type Element interface {
	element()
}

func (e *ModuleDiv) cardModule() string {
	return "div"
}

// ModuleHR 分割线模块
type ModuleHR struct {
	Tag string `json:"tag"` // hr
}

func (e *ModuleHR) cardModule() string {
	return "hr"
}

// ModuleImg 图片模块
type ModuleImg struct {
	Tag     string `json:"tag"` // img
	ImgKey  string `json:"img_key"`
	Title   *Text  `json:"title"`
	Mode    string `json:"mode,omitempty"` // 图片显示模式: crop_center：居中裁剪模式, fit_horizontal：平铺模式
	Alt     *Text  `json:"text,omitempty"`
	Preview bool   `json:"preview,omitemtpy"` // 点击后是否放大图片，缺省为true。在配置 card_link 后可设置为false，使用户点击卡片上的图片也能响应card_link链接跳转
}

func (e *ModuleImg) cardModule() string {
	return "img"
}

// ModuleNote 备注模块
type ModuleNote struct {
	Tag      string             `json:"tag"` // note
	Elements []*ElemTextOrImage `json:"elements"`
	// image
}

func (e *ModuleNote) cardModule() string {
	return "note"
}

// ModuleAction 交互模块
type ModuleAction struct {
	Tag     string          `json:"tag"`
	Actions []ActionElement `json:"actions"`
	Layout  string          `json:"layout,omitemtpy"` // bisected 为二等分布局, trisection 为三等分布局, flow 为流式布局元素会按自身大小横向排列并在空间不够的时候折行
}

func (e *ModuleAction) cardModule() string {
	return "action"
}

type ActionElement interface {
	actionElement()
}

type ElemTextOrImage struct {
	*Text
	*ElemImage
}

// ElemImage 图像元素，注与图像模块的区别（ElementImg)
type ElemImage struct {
	Tag     string `json:"tag"` // tag=img
	ImgKey  string `json:"img_key"`
	Alt     *Text  `json:"alt"`
	Preview bool   `json:"preview,omitemtpy"`
}

func (e *ElemImage) element() {
}

type ElemButton struct {
	Tag      string                 `json:"tag"`                 // tag=button
	Text     *Text                  `json:"text"`                // 按钮中的文本
	URL      string                 `json:"url,omitempty"`       // 跳转链接，和multi_url互斥
	MultiURL *MultiURL              `json:"multi_url,omitempty"` // 多端跳转链接
	Type     string                 `json:"type,omitempty"`      // 配置按钮样式，默认为"default", "default"/"primary"/"danger"
	Value    map[string]interface{} `json:"value,omitempty"`     // 点击后返回业务方
	Confirm  *Confirm               `json:"confirm,omitempty"`   // 二次确认的弹框
}

func (e *ElemButton) element() {
}

func (e *ElemButton) actionElement() {
}

type ElemSelectMenu struct {
	Tag           string                 `json:"tag"` // "select_static" / "select_person", 元素标签，选项模式/选人模式
	PlaceHolder   *Text                  `json:"place_holder,omitempty"`
	InitialOption string                 `json:"initial_option,omitempty"`
	Options       []*Option              `json:"option,omitempty"`
	Value         map[string]interface{} `json:"value,omitempty"`
	Confirm       *Confirm               `json:"confirm,omitempty"`
}

func (e *ElemSelectMenu) element() {
}

func (e *ElemSelectMenu) actionElement() {
}

type ElemOverflow struct {
	Tag     string                 `json:"tag"` // tag=overflow
	Options []*Option              `json:"option"`
	Value   map[string]interface{} `json:"value,omitempty"`
	Confirm *Confirm               `json:"confirm,omitempty"`
}

func (e *ElemOverflow) element() {
}

func (e *ElemOverflow) actionElement() {
}

type ElemDatePicker struct {
	Tag             string                 `json:"tag"` // 如下三种取值 "date_picker", "picker_time", "picker_datetime"
	InitialDate     string                 `json:"initial_date,omitempty"`
	InitialTime     string                 `json:"initial_time,omitempty"`
	InitialDateTime string                 `json:"initial_datetime,omitempty"`
	PlaceHolder     *Text                  `json:"placeholder,omitempty"`
	Value           map[string]interface{} `json:"value,omitempty"`
	Confirm         *Confirm               `json:"confirm,omitempty"`
}

func (e *ElemDatePicker) element() {
}

func (e *ElemDatePicker) actionElement() {
}

// Text 文本对象
type Text struct {
	// Tag 支持"plain_text"和"lark_md"两种模式
	// https://open.feishu.cn/document/ukTMukTMukTM/uADOwUjLwgDM14CM4ATN
	Tag     string `json:"tag"`
	Content string `json:"content"`         // 文本内容
	Lines   int    `json:"lines,omitempty"` // 内容显示行数, 1 显示行数， lines字段仅支持"plain_text"模式
	I18N    *I18N  `json:"i18n,omitempty"`
}

type I18N struct {
	ZHCN string `json:"zh_cn"`
	ENUS string `json:"en_us"`
	JAJP string `json:"jn_jp"`
}

type Field struct {
	IsShort bool  `json:"is_short"`
	Text    *Text `json:"text"`
}

type MultiURL struct {
	URL        string `json:"url"`
	AndroidURL string `json:"android_url"`
	IOSURL     string `json:"ios_url"`
	PCURL      string `json:"pc_url"`
}

type Option struct {
	Value    string    `json:"value"`          // 选项选中后返回业务方的数据
	Text     *Text     `json:"text,omitempty"` // 选项显示内容，非待选人员时必填
	URL      string    `json:"url,omitempty"`  // *仅支持overflow，跳转指定链接，和multi_url字段互斥
	MultiURL *MultiURL `json:"multiURL,omitempty"`
}

type Confirm struct {
	Title *Text `json:"title"` // 弹框标题, 仅支持"plain_text"
	Text  *Text `json:"text"`  // 弹框内容, 仅支持"plain_text"
}

func NewMsgCard(card *Card) *Msg {
	return &Msg{
		MsgType: MsgTypeInteractive,
		Card:    card,
	}
}

func NewMsgInteractiveFromPayload(payload *models.Payload) *Msg {
	// Todo, construct card from payload
	card := &Card{}

	return NewMsgCard(card)
}
