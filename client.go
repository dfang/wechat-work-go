// Package wechatwork 企业微信api的封装
//
// https://work.weixin.qq.com/api/doc#90000/90003/90556
package wechatwork

import (
	"fmt"
	"net/url"
	"os"
	"sync"

	"gopkg.in/resty.v1"
)

// WechatWork 企业微信客户端
type WechatWork struct {
	// CorpID 企业 ID，必填
	CorpID string
}

// App 企业微信客户端（分应用）
type App struct {
	*WechatWork

	// CorpSecret 应用的凭证密钥，其实应该叫AgentSecret更好，必填
	CorpSecret string

	// AgentID 应用 ID，必填
	AgentID int64

	tokenMu     *sync.RWMutex
	AccessToken string
	// tokenExpiresIn time.Duration
	// lastRefresh    time.Time

	// refreshTokenRequestChan chan string // chan currentToken
	// refreshTokenResponseChan chan refreshTokenResult // chan {token, err}

	// Token     string `json:"access_token"`
	ExpiresIn int `json:"expires_in"`
}

// New 构造一个 WechatWork 对象，需要提供企业 ID
func New(corpID string) *WechatWork {
	return &WechatWork{
		CorpID: corpID,
	}
}

// WithApp 构造本企业下某自建 app 的对象
func (app *WechatWork) WithApp(corpSecret string, agentID int64) *App {
	return &App{
		WechatWork: app,

		CorpSecret: corpSecret,
		AgentID:    agentID,

		tokenMu:     &sync.RWMutex{},
		AccessToken: "",
		// lastRefresh: time.Time{},
	}
}

// NewDefaultRestyClient 返回一个resty 的client
func NewDefaultRestyClient() *resty.Client {
	client := resty.New()
	client.SetDebug(true)
	client.SetLogger(os.Stdout)
	client.SetHostURL("https://qyapi.weixin.qq.com")
	return client
}

// NewRequest return resty.Request with right url, right configuration
func (app *App) NewRequest(path string, qs urlValuer, withAccessToken bool) *resty.Request {
	client := NewDefaultRestyClient()

	values := url.Values{}
	if valuer, ok := qs.(urlValuer); ok {
		values = valuer.IntoURLValues()
	}

	if withAccessToken {
		app.SpawnAccessTokenRefresher()
		if app.AccessToken != "" {
			if values.Get("access_token") != "" {
				values.Set("access_token", app.AccessToken)
			} else {
				values.Add("access_token", app.AccessToken)
			}
		}
	}

	url := path + "?" + values.Encode()
	req := client.NewRequest()
	req.URL = url
	return req
}

// Get 一切get请求的api调用可使用此方法
//
// 企业微信中，获取操作和删除都是GET请求
func (app *App) Get(path string, qs urlValuer, respObj interface{}, withAccessToken bool) error {
	req := app.NewRequest(path, qs, withAccessToken)
	resp, err := req.SetResult(&respObj).Get(req.URL)
	if err != nil {
		fmt.Fprintln(os.Stdout, resp.Body())
		panic(err)
	}
	return nil
}

// Post 一切Post请求的api调用使用此方法
//
// 企业微信中，删除操作一般都是GET请求，更新操作、批量删除成员是POST请求，没有PUT、PATCH、DELETE
func (app *App) Post(path string, qs urlValuer, body bodyer, respObj interface{}, withAccessToken bool) error {
	b, _ := body.IntoBody()
	// TODO
	req := app.NewRequest(path, qs, withAccessToken)
	resp, err := req.
		SetHeader("Content-Type", "application/json").
		SetBody(b).
		SetResult(&respObj).
		Post(req.URL)

	if err != nil {
		fmt.Fprintln(os.Stdout, resp.Body())
		panic(err)
	}
	return nil
}
