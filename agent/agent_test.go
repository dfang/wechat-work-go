package agent_test

import (
	"fmt"
	"os"
	"strconv"

	wechatwork "github.com/dfang/wechat-work-go"
	"github.com/dfang/wechat-work-go/agent"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Agent", func() {

	var a *agent.Agent
	// var menu agent.Menu
	var agentID int64
	BeforeEach(func() {})

	JustBeforeEach(func() {
		corpID := os.Getenv("CORP_ID")
		corpSecret := os.Getenv("CORP_SECRET")
		agentID, _ = strconv.ParseInt(os.Getenv("AGENT_ID"), 10, 64)
		client := wechatwork.New(corpID)
		app := client.NewApp(corpSecret, agentID)
		a = agent.WithApp(app)
	})

	Context("应用管理", func() {

		It("获取应用详情", func() {
			result, _ := a.GetAgent(agentID)

			Expect(result.ErrCode).To(Equal(0))
			Expect(result.AgentID).To(Equal(agentID))
		})

		It("获取应用列表", func() {
			result, _ := a.ListAgents()
			fmt.Println(result.AgentList)
			Expect(result.ErrCode).To(Equal(0))
			Expect(len(result.AgentList)).To(BeNumerically(">", 0))
		})

		It("设置应用", func() {
			data := agent.ReqAgentSet{
				AgentID:       agentID,
				IsReportEnter: 1,
			}
			result, _ := a.SetAgent(data)
			Expect(result.ErrCode).To(Equal(0))
		})
	})

	Context("菜单管理", func() {
		var menu agent.Menu
		var respMenuCreate agent.RespCommon
		BeforeEach(func() {
			menu = agent.Menu{
				Button: []agent.Button{
					agent.Button{
						Type: "click",
						Name: "今日歌曲",
						Key:  "V1001_TODAY_MUSIC",
					},
					agent.Button{
						Name: "菜单1",
						SubButton: []agent.SubButton{
							agent.SubButton{
								Name: "菜单1",
								Type: "view",
								URL:  "http://baidu.com",
							},
							agent.SubButton{
								Name: "菜单2",
								Type: "click",
								Key:  "V1001_GOOD",
							},
						},
					},
					agent.Button{
						Type: "click",
						Name: "今日美食",
						Key:  "V1001_TODAY_MUSIC",
					},
				},
			}
		})

		JustBeforeEach(func() {
			// json.NewEncoder(os.Stdout).Encode(menu)
			By("Create menu")
			a.CreateMenu(agentID, menu)
		})

		It("创建菜单", func() {
			Expect(respMenuCreate.ErrCode).To(Equal(0))
		})

		It("获取菜单", func() {
			result, _ := a.GetMenu(agentID)
			Expect(len(result.Button)).To(BeNumerically(">", 0))
		})

		It("删除菜单", func() {
			result, _ := a.DeleteMenu(agentID)
			Expect(result.ErrCode).To(Equal(0))
		})
	})
})
