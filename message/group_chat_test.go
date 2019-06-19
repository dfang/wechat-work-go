package message_test

import (
	"os"
	"strconv"

	wechatwork "github.com/dfang/wechat-work-go"
	"github.com/dfang/wechat-work-go/contact"
	"github.com/dfang/wechat-work-go/message"
	"github.com/dfang/wechat-work-go/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GroupChat", func() {
	var g message.GroupChat
	var c contact.Contact
	var chatid string
	BeforeEach(func() {
		corpID := os.Getenv("CORP_ID")
		corpSecret := os.Getenv("CORP_SECRET")
		agentID, _ := strconv.ParseInt(os.Getenv("AGENT_ID"), 10, 64)

		client := wechatwork.New(corpID)
		app := client.WithApp(corpSecret, agentID)
		g = message.GroupChat{
			App: app,
		}
		contactAppSecret := os.Getenv("CONTACT_APP_SECRET")
		// 关于创建成员（客服答复）
		// 目前只能使用通讯录的secret 获取token进行创建  其他的secret是没有创建成员的权限的
		// 获取路径：通讯录管理secret。在“管理工具”-“通讯录同步”里面查看（需开启“API接口同步”）
		app2 := client.WithApp(contactAppSecret, 0)
		c = contact.Contact{
			App: app2,
		}
	})

	Context("群聊会话", func() {

		It("创建群聊会话", func() {
			var d = models.ReqCreateDepartment{
				Name:     "测试部门",
				ParentID: 1,
				Order:    1,
				ID:       9999,
			}
			c.CreateDepartment(d)

			var u1 = models.ReqMemberCreate{
				UserID:     "zhangsan",
				Name:       "张三",
				Department: []int{9999},
				Mobile:     "12345678901",
			}
			var u2 = models.ReqMemberCreate{
				UserID:     "lisi",
				Name:       "李四",
				Department: []int{9999},
				Mobile:     "15618903010",
			}
			c.CreateMember(u1)
			c.CreateMember(u2)

			req := message.ReqCreateGroupChat{
				Name:     "test",
				UserList: []string{u1.UserID, u2.UserID},
			}

			result, _ := g.CreateGroupChat(req)
			chatid = result.ChatID
			Expect(result.ErrCode).To(Equal(0))
		})

		It("修改群聊会话", func() {
			req := message.ReqUpdateGroupChat{
				ChatID: chatid,
				Name:   "hello wechat work",
			}
			result, _ := g.UpdateGroupChat(req)
			Expect(result.ErrCode).To(Equal(0))
		})

		It("获取群聊会话", func() {
			result, _ := g.GetGroupChat(chatid)
			Expect(result.ErrCode).To(Equal(0))
			Expect(result.ChatInfo.ChatID).To(Equal(chatid))
		})

	})
})
