package WechatWork_test

import (
	"os"
	"strconv"

	WechatWork "github.com/dfang/wechat-work-go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Access Token", func() {
	corpID := os.Getenv("CORP_ID")
	corpSecret := os.Getenv("CORP_SECRET")
	agentID, _ := strconv.ParseInt(os.Getenv("AGENT_ID"), 10, 64)
	var app *WechatWork.App

	BeforeEach(func() {
		client := WechatWork.New(corpID)
		app = client.WithApp(corpSecret, agentID)
	})

	It("should get access token", func() {
		app.SpawnAccessTokenRefresher()
		Expect(app.AccessToken).NotTo(BeNil())
	})

	It("运行了SyncAccessToken后, app.AccessToken应该不为空", func() {
		app.SyncAccessToken()
		Expect(app.AccessToken).NotTo(BeNil())
	})

})
