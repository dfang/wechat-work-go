package WechatWork

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"sync"

	"github.com/go-resty/resty"
)

// WechatWork 企业微信客户端
type WechatWork struct {
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

	tokenMu     *sync.RWMutex
	AccessToken string
	// tokenExpiresIn time.Duration
	// lastRefresh    time.Time

	// refreshTokenRequestChan chan string // chan currentToken
	// refreshTokenResponseChan chan refreshTokenResult // chan {token, err}

	// Token     string `json:"access_token"`
	ExpiresIn int64 `json:"expires_in"`
}

// New 构造一个 WechatWork 客户端对象，需要提供企业 ID
func New(corpID string) *WechatWork {
	return &WechatWork{
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
		// lastRefresh: time.Time{},
	}
}

func (c *App) NewRestyClient() *resty.Client {
	client := resty.New()
	// client.SetDebug(true)
	client.SetHostURL("https://qyapi.weixin.qq.com")
	client.SetDebug(true)
	return client
}

// NewRequest return resty.Request with right url
func (c *App) NewRequest(path string, qs urlValuer, withAccessToken bool) *resty.Request {
	client := resty.New()
	client.SetDebug(true)
	// client.SetLogger(os.Stdout)
	client.SetHostURL("https://qyapi.weixin.qq.com")

	values := url.Values{}
	if valuer, ok := qs.(urlValuer); ok {
		values = valuer.IntoURLValues()
	}

	if withAccessToken {
		c.SyncAccessToken()
		// c.SpawnAccessTokenRefresher()
		if c.AccessToken != "" {
			if values.Get("access_token") != "" {
				values.Set("access_token", c.AccessToken)
			} else {
				values.Add("access_token", c.AccessToken)
			}
		}
	}

	url := path + "?" + values.Encode()
	// client.R().URL = url
	req := client.NewRequest()
	req.URL = url
	return req
}

// Get  Get 请求的api调用
func (c *App) Get(path string, qs urlValuer, respObj interface{}, withAccessToken bool) error {
	client := resty.New()
	// client.SetDebug(true)
	client.SetHostURL("https://qyapi.weixin.qq.com")

	values := url.Values{}
	if valuer, ok := qs.(urlValuer); ok {
		values = valuer.IntoURLValues()
	}

	// c.SpawnAccessTokenRefresher()
	// values.Add("access_token", c.AccessToken)

	if withAccessToken {
		c.SyncAccessToken()
		// c.SpawnAccessTokenRefresher()
		if c.AccessToken != "" {
			if values.Get("access_token") != "" {
				values.Set("access_token", c.AccessToken)
			} else {
				values.Add("access_token", c.AccessToken)
			}
		}
	}

	url := path + "?" + values.Encode()
	resp, err := client.R().SetResult(respObj).Get(url)
	if err != nil {
		fmt.Fprintln(os.Stdout, resp.Body())
		panic(err)
	}
	return nil
}

// Post Post 请求的api调用
func (c *App) Post(path string, qs urlValuer, body bodyer, respObj interface{}, withAccessToken bool) (interface{}, error) {
	// url := c.composeQyapiURLWithToken(path, req, withAccessToken)
	// urlStr := url.String()
	client := resty.New()
	client.SetDebug(true)
	client.SetHostURL("https://qyapi.weixin.qq.com")

	values := url.Values{}
	if valuer, ok := qs.(urlValuer); ok {
		values = valuer.IntoURLValues()
	}

	if withAccessToken {
		c.SyncAccessToken()
		// c.SpawnAccessTokenRefresher()
		if c.AccessToken != "" {
			if values.Get("access_token") != "" {
				values.Set("access_token", c.AccessToken)
			} else {
				values.Add("access_token", c.AccessToken)
			}
		}
	}

	url := path + "?" + values.Encode()

	b, _ := body.IntoBody()
	// TODO

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(b).
		SetResult(&respObj).
		Post(url)

	if err != nil {
		panic(err)
	}
	// defer resp.Close()

	decoder := json.NewDecoder(bytes.NewReader(resp.Body()))
	err = decoder.Decode(&respObj)
	if err != nil {
		return respObj, err
	}
	return respObj, nil
}
