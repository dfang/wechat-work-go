package WechatWork

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/go-resty/resty"
)

// WechatWork 企业微信客户端
type WechatWork struct {
	opts options

	// CorpID 企业 ID，必填
	CorpID string
}

// App 企业微信客户端（分应用）
type App struct {
	*WechatWork

	// CorpSecret 应用的凭证密钥，必填
	CorpSecret string
	// AgentID 应用 ID，必填
	AgentID int64

	tokenMu        *sync.RWMutex
	AccessToken    string
	tokenExpiresIn time.Duration
	lastRefresh    time.Time
}

// New 构造一个 WechatWork 客户端对象，需要提供企业 ID
func New(corpID string, opts ...ctorOption) *WechatWork {
	optionsObj := defaultOptions()

	for _, o := range opts {
		o.ApplyTo(&optionsObj)
	}

	return &WechatWork{
		opts: optionsObj,

		CorpID: corpID,
	}
}

// WithApp 构造本企业下某自建 app 的客户端
func (c *WechatWork) WithApp(corpSecret string, agentID int64) *App {
	return &App{
		WechatWork: c,

		CorpSecret: corpSecret,
		AgentID:    agentID,

		tokenMu:     &sync.RWMutex{},
		AccessToken: "",
		lastRefresh: time.Time{},
	}
}

// Get Get 请求
func (c *App) Get(path string, req urlValuer, respObj interface{}, withAccessToken bool) error {
	// url := c.composeQyapiURLWithToken(path, req, withAccessToken)
	// urlStr := url.String()

	client := resty.New()
	client.SetDebug(true)
	client.SetHostURL("https://qyapi.weixin.qq.com")

	values := url.Values{}
	if valuer, ok := req.(urlValuer); ok {
		values = valuer.IntoURLValues()
	}

	// c.SpawnAccessTokenRefresher()
	// values.Add("access_token", c.AccessToken)

	if withAccessToken {
		c.SyncAccessToken()
		// c.SpawnAccessTokenRefresher()

		if c.AccessToken != "" {
			values.Add("access_token", c.AccessToken)
		}

		// fmt.Println(c.AccessToken)
	}

	url := path + "?" + values.Encode()

	// fmt.Println(url)

	resp, err := client.R().Get(url)
	if err != nil {
		// panic("err when requesting api request to qyapi.weixin.qq.com")
		panic(err)
	}
	// defer resp.Close()

	decoder := json.NewDecoder(bytes.NewReader(resp.Body()))
	err = decoder.Decode(&respObj)
	if err != nil {
		return err
	}
	return nil
}

//
// impl App
//

func (c *App) composeQyapiURL(path string, req interface{}) *url.URL {
	values := url.Values{}
	if valuer, ok := req.(urlValuer); ok {
		values = valuer.IntoURLValues()
	}

	// TODO: refactor
	base, err := url.Parse(c.opts.QYAPIHost)
	if err != nil {
		// TODO: error_chain
		panic(fmt.Sprintf("qyapiHost invalid: host=%s err=%+v", c.opts.QYAPIHost, err))
	}

	base.Path = path
	base.RawQuery = values.Encode()

	return base
}

func (c *App) composeQyapiURLWithToken(path string, req interface{}, withAccessToken bool) *url.URL {
	url := c.composeQyapiURL(path, req)

	if !withAccessToken {
		return url
	}

	// intensive mutex juggling action
	c.tokenMu.RLock()
	if c.AccessToken == "" {
		c.tokenMu.RUnlock() // RWMutex doesn't like recursive locking
		// TODO: what to do with the possible error?
		_ = c.SyncAccessToken()
		c.tokenMu.RLock()
	}
	tokenToUse := c.AccessToken
	c.tokenMu.RUnlock()

	q := url.Query()
	q.Set("access_token", tokenToUse)
	url.RawQuery = q.Encode()

	return url
}

func (c *App) executeQyapiGet(path string, req urlValuer, respObj interface{}, withAccessToken bool) error {
	url := c.composeQyapiURLWithToken(path, req, withAccessToken)
	urlStr := url.String()

	// fmt.Println(url)
	// fmt.Println(urlStr)

	// resp, err := resty.R().Get("http://httpbin.org/get")

	resp, err := c.opts.HTTP.Get(urlStr)
	if err != nil {
		// TODO: error_chain
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(respObj)
	if err != nil {
		// TODO: error_chain
		return err
	}

	return nil
}

func (c *App) executeQyapiJSONPost(path string, req bodyer, respObj interface{}, withAccessToken bool) error {
	url := c.composeQyapiURLWithToken(path, req, withAccessToken)
	urlStr := url.String()

	body, err := req.IntoBody()
	if err != nil {
		// TODO: error_chain
		return err
	}

	resp, err := c.opts.HTTP.Post(urlStr, "application/json", bytes.NewReader(body))
	if err != nil {
		// TODO: error_chain
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(respObj)
	if err != nil {
		// TODO: error_chain
		return err
	}

	return nil
}
