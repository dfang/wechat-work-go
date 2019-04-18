package WechatWork_test

import (
	"os"
	"strconv"
	"testing"

	WechatWork "github.com/dfang/wechat-work-go"
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
func TestExampleWechatWork(t *testing.T) {
	// app, fn := setupTest(t)
	// app.SyncAccessToken()

	// corpID := os.Getenv("CORP_ID")
	// corpSecret := os.Getenv("CORP_SECRET")
	// agentID, _ := strconv.ParseInt(os.Getenv("AGENT_ID"), 10, 64)

	// client := WechatWork.New(corpID)

	// there're advanced options
	// _ = New(
	// 	corpID,
	// 	WithQYAPIHost("http://localhost:8888"),
	// 	WithHTTPClient(&http.Client{}),
	// )

	// work with individual apps
	// app := client.WithApp(corpSecret, agentID)
	// app.SpawnAccessTokenRefresher()

	// t.Log("AccessToken is ")
	// t.Log(app.AccessToken)
	// see other examples for more details
}

func TestExampleApp_SendTextMessage(t *testing.T) {
	// corpID := os.Getenv("CORP_ID")
	// corpSecret := os.Getenv("CORP_SECRET")
	// agentID, _ := strconv.ParseInt(os.Getenv("AGENT_ID"), 10, 64)

	// client := WechatWork.New(corpID)

	// app := client.WithApp(corpSecret, agentID)
	// // preferably do this at app initialization
	// app.SpawnAccessTokenRefresher()

	app, fn := setupTest(t)
	app.SyncAccessToken()
	fn()
	// app.SyncAccessToken()
	// t.Log("AccessToken is ")
	// t.Log(app.AccessToken)

	// send to user(s)
	to1 := WechatWork.Recipient{
		UserIDs: []string{"fang"},
	}
	_ = app.SendTextMessage(&to1, "啊啊啊啊send to user(s)", false)

	// t.Log("AccessToken is ")
	// t.Log(app.AccessToken)

	// "safe" message
	to2 := WechatWork.Recipient{
		UserIDs: []string{"fang"},
	}
	_ = app.SendTextMessage(&to2, "safe message", true)

	// send to party(parties)
	to3 := WechatWork.Recipient{
		PartyIDs: []string{"testdept"},
	}
	_ = app.SendTextMessage(&to3, "send to party(parties)", false)

	// send to tag(s)
	to4 := WechatWork.Recipient{
		TagIDs: []string{"testtag"},
	}
	_ = app.SendTextMessage(&to4, "send to tag(s)", false)

	// send to chatid
	to5 := WechatWork.Recipient{
		ChatID: "CHATID",
	}
	_ = app.SendTextMessage(&to5, "send to chatid", false)

	// to6 := Recipient{}
}
