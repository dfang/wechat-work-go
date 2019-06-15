package WechatWork_test

import (
	"os"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	WechatWork "github.com/dfang/wechat-work-go"
)

var _ = Describe("Agent", func() {
	corpID := os.Getenv("CORP_ID")
	corpSecret := os.Getenv("CORP_SECRET")
	agentID, _ := strconv.ParseInt(os.Getenv("AGENT_ID"), 10, 64)
	var app *WechatWork.App

	BeforeEach(func() {
		client := WechatWork.New(corpID)
		app = client.WithApp(corpSecret, agentID)
		app.SpawnAccessTokenRefresher()
	})

	Context("应用管理", func() {
		It("获取应用信息", func() {
			resp, _ := app.GetAgent("0")

			// fmt.Println(app.CorpID)
			// fmt.Println(app.CorpSecret)
			// fmt.Println(app.AgentID)

			Expect(resp.AgentID).To(Equal(0))
		})

		It("获取access_token下的应用列表", func() {
			resp, _ := app.ListAgents()
			Expect(resp.ErrCode).To(Equal(0))
			Expect(resp.ErrMsg).To(Equal("ok"))
		})
	})
})
