package WechatWork_test

import (
	"os"
	"strconv"
	"testing"

	WechatWork "github.com/dfang/wechat-work-go"
	c "github.com/smartystreets/goconvey/convey"
)

func setupTest(t *testing.T) (*WechatWork.App, func()) {
	// Test setup
	t.Log("setupTest()")

	corpID := os.Getenv("CORP_ID")
	corpSecret := os.Getenv("CORP_SECRET")
	agentID, _ := strconv.ParseInt(os.Getenv("AGENT_ID"), 10, 64)

	client := WechatWork.New(corpID)
	app := client.WithApp(corpSecret, agentID)
	// preferably do this at app initialization
	app.SpawnAccessTokenRefresher()

	// Test teardown - return a closure for use by 'defer'
	return app, func() {
		// t is from the outer setupTest scope
		t.Log("teardownTest()")
	}
}

func TestWechatWorkGetAccessToken(t *testing.T) {
	// app, fn := defer setupTest(t)()
	app, fn := setupTest(t)
	app.SyncAccessToken()

	c.Convey("获取AccessToken", t, func() {
		c.Convey("运行了SyncAccessToken后 app.AccessToken应该不为空", func() {
			// c.Printf("\nAccessToken: %s\n", app.AccessToken)
			c.So(app.AccessToken, c.ShouldNotEqual, "")
		})

		fn()
	})
}
