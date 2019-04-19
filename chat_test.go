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
	app.SyncAccessToken()

	// Test teardown - return a closure for use by 'defer'
	return app, func() {
		// t is from the outer setupTest scope
		t.Log("teardownTest()")
	}
}

func TestWechatWorkCreateAppChat(t *testing.T) {
	// app, fn := defer setupTest(t)()
	app, fn := setupTest(t)
	app.SyncAccessToken()

	c.Convey("创建群聊会话", t, func() {
		a := WechatWork.ReqAppChatCreate{
			Name:     "TEST",
			Owner:    "fang",
			Userlist: []string{"fang", "test01"},
			Chatid:   "TEST",
		}
		r, _ := app.CreateAppChat(a)
		t.Logf("\n%+v\n", r)

		c.So(r.CHATID, c.ShouldEqual, "TEST")

		fn()
	})
}

func TestWechatWorkUpdateAppChat(t *testing.T) {
	// app, fn := defer setupTest(t)()
	app, fn := setupTest(t)
	app.SyncAccessToken()

	c.Convey("修改群聊会话", t, func() {
		a := WechatWork.ReqAppChatUpdate{
			Chatid:      "TEST",
			Name:        "TEST",
			Owner:       "fang",
			AddUserList: []string{"fang", "test02", "test03"},
			DelUserList: []string{"test01"},
		}
		r, _ := app.UpdateAppChat(a)
		t.Logf("\n%+v\n", r)

		c.So(r.ErrCode, c.ShouldEqual, 0)
		c.So(r.ErrMsg, c.ShouldEqual, "ok")

		fn()
	})
}

func TestWechatWorkGetAppChat(t *testing.T) {
	app, fn := setupTest(t)

	c.Convey("获取群聊会话", t, func() {
		r, _ := app.GetAppChat("TEST")
		t.Logf("\n%+v\n", r)

		c.So(r.ErrCode, c.ShouldEqual, 0)
		c.So(r.ErrMsg, c.ShouldEqual, "ok")
		c.So(r.ChatInfo.Name, c.ShouldEqual, "TEST")
		c.So(r.ChatInfo.Chatid, c.ShouldEqual, "TEST")
		fn()
	})
}

func TestWechatWorkSendAppChatMessage(t *testing.T) {
	app, fn := setupTest(t)
	c.Convey("发送文本消息", t, func() {
		// req := WechatWork.AppChatTextMessage{}
		// req.AppChatMessage.Chatid = "TEST"
		// req.AppChatMessage.Msgtype = "text"
		// req.Text.Content = "xxxxx"
		// req.Safe = 0

		req := WechatWork.AppChatTextCardMessage{}
		req.AppChatMessage.Chatid = "TEST"
		req.AppChatMessage.Msgtype = "textcard"
		req.Safe = 0
		req.Textcard = struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			URL         string `json:"url"`
			Btntxt      string `json:"btntxt"`
		}{
			Title:       "this is title",
			Description: "this is description",
			URL:         "http://baidu.com",
			Btntxt:      "btn",
		}
		t.Logf("\n%+v\n", req)

		r, _ := app.SendAppChat(req)
		t.Logf("\n%+v\n", r)

		c.So(r.ErrCode, c.ShouldEqual, 0)
		c.So(r.ErrMsg, c.ShouldEqual, "ok")
		// c.So(r.ChatInfo.Name, c.ShouldEqual, "TEST")
		// c.So(r.ChatInfo.Chatid, c.ShouldEqual, "TEST")
		fn()
	})
}
