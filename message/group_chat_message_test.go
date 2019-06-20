package message_test

import (
	"github.com/dfang/wechat-work-go/message"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GroupChat", func() {
	var chatid string
	BeforeEach(func() {})

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
