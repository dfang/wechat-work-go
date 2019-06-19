package agent_test

import (
	"os"
	"strconv"

	wechatwork "github.com/dfang/wechat-work-go"
	"github.com/dfang/wechat-work-go/agent"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Agent", func() {

	var a *agent.Agent

	BeforeEach(func() {
		corpID := os.Getenv("CORP_ID")
		corpSecret := os.Getenv("CORP_SECRET")
		agentID, _ := strconv.ParseInt(os.Getenv("AGENT_ID"), 10, 64)

		client := wechatwork.New(corpID)
		app := client.WithApp(corpSecret, agentID)
		a = &agent.Agent{
			App: app,
		}
	})

	Context("应用管理", func() {

		It("获取应用详情", func() {
			result, _ := a.GetAgent("0")

			Expect(result.ErrCode).To(Equal(0))
			Expect(result.AgentID).To(Equal(0))
		})

		// It("获取应用列表", func() {
		// 	result, _ := a.ListAgents()
		// 	Expect(result.ErrCode).To(Equal(0))
		// 	Expect(len(result.AgentList)).To(BeNumerically(">", 0))
		// })

		// It("设置应用", func() {
		// 	data := models.ReqAgentSet{
		// 		AgentID:       "0",
		// 		IsReportEnter: 1,
		// 	}
		// 	result, _ := a.SetAgent(data)
		// 	Expect(result.ErrCode).To(Equal(0))

		// })

		// It("创建菜单", func() {
		// 	menu := models.Menu{
		// 		Button: []models.Button{
		// 			models.Button{
		// 				Type: "click",
		// 				Name: "今日歌曲",
		// 				Key:  "V1001_TODAY_MUSIC",
		// 			},
		// 			models.Button{
		// 				Name: "菜单1",
		// 				SubButton: []models.SubButton{
		// 					models.SubButton{
		// 						Name: "菜单1",
		// 						Type: "view",
		// 						URL:  "http://baidu.com",
		// 					},
		// 					models.SubButton{
		// 						Name: "菜单2",
		// 						Type: "click",
		// 						Key:  "V1001_GOOD",
		// 					},
		// 				},
		// 			},
		// 			models.Button{
		// 				Type: "click",
		// 				Name: "今日美食",
		// 				Key:  "V1001_TODAY_MUSIC",
		// 			},
		// 		},
		// 	}

		// 	json.NewEncoder(os.Stdout).Encode(menu)
		// 	result, _ := m.CreateMenu("0", menu)
		// 	Expect(result.ErrCode).To(Equal(0))
		// })

		// It("获取菜单", func() {
		// 	result, _ := m.GetMenu("0")
		// 	Expect(len(result.Button)).To(BeNumerically(">", 0))
		// })

		// It("删除菜单", func() {
		// 	result, _ := m.DeleteMenu("0")
		// 	Expect(result.ErrCode).To(Equal(0))
		// })
	})
})
