package weixinapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bougou/webhook-adapter/models"
)

// 企业微信 - 应用

const ChannelTypeWeixin = "weixinapp"

type Payload2Msg func(payload *models.Payload) *Msg

var SupportedMsgtypes = make(map[string]Payload2Msg)

func ValidMsgtype(msgtype string) bool {
	if _, exists := SupportedMsgtypes[msgtype]; !exists {
		return false
	}
	return true
}

// ErrCodeAboutTokens contains weixinapp
// https://work.weixin.qq.com/api/doc/90000/90139/90313
var ErrCodeAboutTokens = []int{
	40014, // 不合法的access_token
	42001, // access_token已过期
}

type Notifier struct {
	addr        string
	corpID      string // 企业ID
	agentID     int    // 应用ID
	agentSecret string // 应用的凭证密钥, 区分应用
	client      *http.Client

	token          string
	tokenAt        time.Time
	tokenExpiredIn time.Duration
}

func NewNotifer(corpID string, agentID int, agentSecret string) *Notifier {
	return &Notifier{
		addr:        "https://qyapi.weixin.qq.com",
		corpID:      corpID,
		agentID:     agentID,
		agentSecret: agentSecret,
		client:      &http.Client{},
	}
}

func (n *Notifier) Addr() string {
	return fmt.Sprintf("%s/cgi-bin/message/send?access_token=%s", n.addr, n.token)
}

func (n *Notifier) AddrForGetToken() string {
	return fmt.Sprintf("%s/cgi-bin/gettoken?corpid=%s&corpsecret=%s", n.addr, n.corpID, n.agentSecret)
}

func (n *Notifier) GetToken() error {
	req, err := http.NewRequest("GET", n.AddrForGetToken(), nil)
	if err != nil {
		return fmt.Errorf("get token construct http request failed, %s", err)
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := n.client.Do(req)
	if err != nil {
		return fmt.Errorf("get token request failed, %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("get token response error, status: %d", res.StatusCode)
	}

	type ResponseBody struct {
		ErrCode          int    `json:"errcode"`
		ErrMsg           string `json:"errmsg"`
		AccessToken      string `json:"access_token"`
		ExpiresInSeconds int    `json:"expires_in"`
	}

	r := &ResponseBody{}
	if err := json.NewDecoder(res.Body).Decode(r); err != nil {
		return fmt.Errorf("get token decode response body failed, %s", err)
	}

	if r.ErrCode != 0 {
		return fmt.Errorf("get token failed, errmsg is %s", r.ErrMsg)
	}

	n.token = r.AccessToken
	n.tokenAt = time.Now()
	n.tokenExpiredIn, err = time.ParseDuration(fmt.Sprintf("%ds", r.ExpiresInSeconds))
	if err != nil {
		n.tokenExpiredIn = 2 * time.Hour
	}

	return nil

}

func (n *Notifier) ShouldGetToken() bool {
	if n.token == "" || time.Since(n.tokenAt) > n.tokenExpiredIn {
		return true
	}

	return false
}

func (n *Notifier) Send(msg *Msg) error {
	// fill agentID
	msg.AgentID = n.agentID

	if !msg.Valid() {
		msg.ToUser = "@all"
	}

	if n.token == "" {
		if err := n.GetToken(); err != nil {
			return fmt.Errorf("failed token failed, %s", err)
		}
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", n.Addr(), bytes.NewBuffer(msgBytes))
	if err != nil {
		return fmt.Errorf("failed to construct request, got %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := n.client.Do(req)
	if err != nil {
		return fmt.Errorf("send msg error, %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("send msg response error, status: %d", res.StatusCode)
	}

	return nil
}
