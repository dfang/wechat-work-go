package WechatWork

import (
	"time"

	"github.com/cenkalti/backoff"
)

// getAccessToken 获取 access token
func (c *App) getAccessToken() (RespAccessToken, error) {
	req := reqAccessToken{
		CorpID:     c.CorpID,
		CorpSecret: c.CorpSecret,
	}

	var resp RespAccessToken
	err := c.executeQyapiGet("/cgi-bin/gettoken", req, &resp, false)
	if err != nil {
		// TODO: error_chain
		return RespAccessToken{}, err
	}

	return resp, nil
}

// SyncAccessToken 同步该客户端实例的 access token
//
// 会拿 `tokenMu` 写锁
func (c *App) SyncAccessToken() error {
	tok, err := c.getAccessToken()
	if err != nil {
		// TODO: error_chain
		return err
	}

	c.tokenMu.Lock()
	defer c.tokenMu.Unlock()

	c.AccessToken = tok.AccessToken
	c.tokenExpiresIn = time.Duration(tok.ExpiresInSecs) * time.Second
	c.lastRefresh = time.Now()

	return nil
}

func (c *App) AccessTokenRefresher() {
	const refreshTimeWindow = 30 * time.Minute
	const minRefreshDuration = 5 * time.Second

	// TODO: context cancellation
	for {
		retryer := backoff.NewExponentialBackOff()
		err := backoff.Retry(c.SyncAccessToken, retryer)
		if err != nil {
			// wtf
			// TODO: logging
			_ = err
		}

		waitUntilTime := c.lastRefresh.Add(c.tokenExpiresIn).Add(-refreshTimeWindow)
		waitDuration := time.Until(waitUntilTime)

		if waitDuration < minRefreshDuration {
			waitDuration = minRefreshDuration
		}

		time.Sleep(waitDuration)
	}
}

// SpawnAccessTokenRefresher 启动该 app 的 access token 刷新 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *App) SpawnAccessTokenRefresher() {
	go c.AccessTokenRefresher()
}
