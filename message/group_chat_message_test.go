package message_test

import (
	"github.com/dfang/wechat-work-go/contact"
	"github.com/dfang/wechat-work-go/message"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var userIDForChat1 string
var userIDForChat2 string

var _ = Describe("GroupChat", func() {
	var chatid string
	BeforeEach(func() {
		userIDForChat1 = "chat1"
		userIDForChat2 = "chat2"
		chatid = "groupchat_test"
	})
	JustBeforeEach(func() {
		createTestDepartmentForChatTest()
		createTestUsersForChat()
	})

	// TODO
	// 群聊会话 不支持通过 API 来删除
	// 对于完整的测试很不利
	// CI 经常会出错
	Context("群聊会话", func() {

		var res1 message.RespCreateGroupChat
		var res2 message.RespCommon

		JustBeforeEach(func() {
			req1 := message.ReqCreateGroupChat{
				Name:     chatid,
				UserList: []string{userIDForChat1, userIDForChat2},
			}
			res1, _ = g.CreateGroupChat(req1)

			req2 := message.ReqUpdateGroupChat{
				ChatID: chatid,
				Name:   "hello wechat work",
			}
			res2, _ = g.UpdateGroupChat(req2)
		})

		It("创建群聊会话", func() {
			// Expect(res1.ErrCode).To(Equal(0))
			Expect(res1.ChatID).To(Equal("groupchat_test"))
		})

		It("修改群聊会话", func() {
			Expect(res2.ErrCode).To(Equal(0))
		})

		It("获取群聊会话", func() {
			result, _ := g.GetGroupChat(chatid)
			Expect(result.ErrCode).To(Equal(0))
			Expect(result.ChatInfo.ChatID).To(Equal(chatid))
		})

	})
})

func createTestDepartmentForChatTest() {
	var d = contact.ReqCreateDepartment{
		Name:     "群聊会话测试部门",
		ParentID: 1,
		Order:    1,
		ID:       9999,
	}
	_, _ = c.CreateDepartment(d)
}

func createTestUsersForChat() {
	var u1 = contact.ReqMemberCreate{
		UserID:     "chat1",
		Name:       "张三",
		Department: []int{9999},
		Mobile:     "12345678901",
	}
	var u2 = contact.ReqMemberCreate{
		UserID:     "chat2",
		Name:       "李四",
		Department: []int{9999},
		Mobile:     "15618903010",
	}
	_, _ = c.CreateMember(u1)
	_, _ = c.CreateMember(u2)
}
