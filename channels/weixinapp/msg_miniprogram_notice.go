package weixinapp

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
