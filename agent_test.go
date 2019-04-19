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

func TestWechatWorkGetAgent(t *testing.T) {
	// app, fn := defer setupTest(t)()
	app, fn := setupTest(t)

	c.Convey("获取应用信息", t, func() {
		r, err := app.GetAgent(1000002)
		t.Logf("response: %+v", r)

		c.Convey("应用信息 应该正确获取了", func() {
			c.So(err, c.ShouldEqual, nil)
			c.So(r.Agentid, c.ShouldEqual, 1000002)
		})

		fn()
	})
}

func TestWechatWorkSetAgent(t *testing.T) {
	app, fn := setupTest(t)

	c.Convey("设置应用信息", t, func() {
		req := WechatWork.ReqAgentSet{
			Agentid:            1000002,
			ReportLocationFlag: 1,
			Name:               "Test",
			Description:        "内部频道",
			Isreportenter:      1,
		}
		r, err := app.SetAgent(req)
		t.Logf("response: %+v", r)

		c.Convey("应用信息 应该正确设置了", func() {
			c.So(err, c.ShouldEqual, nil)
			c.So(r.ErrCode, c.ShouldEqual, 0)
			c.So(r.ErrMsg, c.ShouldEqual, "ok")
			// c.So(r.Name, c.ShouldEqual, "Test")
		})

		fn()
	})
}

func TestWechatWorkCreateMenu(t *testing.T) {
	// app, fn := defer setupTest(t)()
	app, fn := setupTest(t)

	c.Convey("获取应用信息", t, func() {
		r, err := app.GetAgent(1000002)
		t.Logf("response: %+v", r)

		c.Convey("应用信息 应该正确获取了", func() {
			c.So(err, c.ShouldEqual, nil)
			c.So(r.Agentid, c.ShouldEqual, 1000002)
		})

		fn()
	})
}
