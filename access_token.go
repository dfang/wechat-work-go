package wechatwork

import (
	"fmt"
	"sync"
	"time"

	"github.com/dfang/wechat-work-go/models"
	"gopkg.in/resty.v1"
)

// getAccessToken 获取 access token
// https://work.weixin.qq.com/api/doc#90000/90135/91039
func (app *App) getAccessToken() (models.RespAccessToken, error) {
	apiPath := "/cgi-bin/gettoken"
	resty.SetHostURL("https://qyapi.weixin.qq.com")
	resty.SetQueryParam("corpid", app.CorpID)
	resty.SetQueryParam("corpsecret", app.CorpSecret)
	resty.SetDebug(true)

	var data models.RespAccessToken
	resp, err := resty.R().SetResult(&data).Get(apiPath)
	if err != nil {
		// log.Println("err when request gettoken api")
		// fmt.Println(err)
		// fmt.Println(resp.Status())
		// fmt.Println(resp.StatusCode())
		// fmt.Printf("%+v\n", data)
		return models.RespAccessToken{}, err
	}

	// 40001 不合法的secret参数
	// 40013 不合法的CorpID
	// 40056 不合法的agentid
	// 前两种情况 直接panic就好了 无需retry
	// access_token是分应用的，获取access_token 无需AgentID
	// 但是操作某些api 需要传AgentID
	if data.ErrCode == 40013 {
		panic("请检查CorpID 参数")
	}

	if data.ErrCode == 40001 {
		panic("请检查CorpSecret 参数")
	}

	// fmt.Println("get token ......")
	// fmt.Println(resp.Status())
	fmt.Println(resp.StatusCode())
	// fmt.Printf("%+v\n", data)

	// 全局错误码 https://work.weixin.qq.com/api/doc#90000/90139/90313
	// -1 表示系统繁忙
	if data.ErrCode == -1 {
		// TODO: retry logic here
	}

	return data, nil
}

// SyncAccessToken 同步该app实例的 access token
func (app *App) SyncAccessToken() error {
	tok, err := app.getAccessToken()
	if err != nil {
		fmt.Println(err)
		return err
	}

	app.tokenMu.Lock()
	defer app.tokenMu.Unlock()

	app.AccessToken = tok.AccessToken
	app.ExpiresIn = tok.ExpiresInSecs
	return nil
}

// AccessTokenRefresher Refresh ExpiresIn in Ticker
func (app *App) AccessTokenRefresher(o *sync.Once) {
	o.Do(func() {
		if app.AccessToken == "" {
			app.SyncAccessToken()
		}
		tickDuration := time.Minute * 1
		ticker := time.NewTicker(tickDuration)
		for {
			select {
			case <-ticker.C:
				// fmt.Println(".........tock .....")
				// fmt.Println("access_token", app.AccessToken)
				// fmt.Println("expires_in", app.ExpiresIn)
				if app.ExpiresIn <= 7140 {
					fmt.Println("expires_in 剩余时间不多，重新获取access_token")
					// 如果ExpiresIn < 10 分钟那就重新发起请求获取
					app.SyncAccessToken()
				} else {
					app.ExpiresIn -= 60
					fmt.Println("after reduce expires_in")
					fmt.Println("access_token", app.AccessToken)
					fmt.Println("expires_in", app.ExpiresIn)
				}
			}
		}
	})
}

// SpawnAccessTokenRefresher 启动该 app 的 access token 刷新 goroutine
func (app *App) SpawnAccessTokenRefresher() {
	once := &sync.Once{}
	go app.AccessTokenRefresher(once)
}
