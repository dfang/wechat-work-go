package message_test

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

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
		chatid = randChatID()
		// chatid = "7kC8cZ5aGgooMPWKkM0tQU8cCBYuCcAI"
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
				ChatID:   chatid,
			}
			res1, _ = msg.CreateGroupChat(req1)

			req2 := message.ReqUpdateGroupChat{
				ChatID:      chatid,
				Name:        "hello wechat work",
				Owner:       *selfID,
				AddUserList: []string{*selfID},
			}
			res2, _ = msg.UpdateGroupChat(req2)
		})

		It("创建群聊会话", func() {
			// Expect(res1.ErrCode).To(Equal(0))
			Expect(res1.ChatID).To(Equal(chatid))
		})

		It("修改群聊会话", func() {
			Expect(res2.ErrCode).To(Equal(0))
		})

		It("获取群聊会话", func() {
			result, _ := msg.GetGroupChat(chatid)
			Expect(result.ErrCode).To(Equal(0))
			Expect(result.ChatInfo.ChatID).To(Equal(chatid))
		})

		It("发送消息到群聊会话", func() {
			msg1 := message.GroupChatTextMessage{}
			msg1.ChatID = chatid
			msg1.Safe = 0
			msg1.Text.Content = "this is a test message sent by api in test"
			msg1.MsgType = "text"
			result, _ := msg.SendGroupChatMessage(msg1)
			Expect(result.ErrCode).To(Equal(0))
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
	var u1 = contact.ReqCreateMember{
		UserID:     "chat1",
		Name:       randCNName(),
		Department: []int{9999},
		Mobile:     randCNPhone(),
		Enable:     1,
	}
	var u2 = contact.ReqCreateMember{
		UserID:     "chat2",
		Name:       randCNName(),
		Department: []int{9999},
		Mobile:     randCNPhone(),
		Enable:     1,
	}
	_, _ = c.CreateMember(u1)
	_, _ = c.CreateMember(u2)
}

func randCNName() string {
	rand.Seed(time.Now().UnixNano())
	var fList []string = []string{"赵", "钱", "孙", "李", "周", "吴", "郑", "王", "冯", "陈", "褚", "卫", "蒋", "沈", "韩", "杨", "朱"}
	var lList []string = []string{"一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}
	f := fList[rand.Intn(len(fList)-1)]
	l := lList[rand.Intn(len(lList)-1)]
	return f + l
}

func randCNPhone() string {
	rand.Seed(time.Now().UnixNano())
	prefix := []string{"130", "131", "132", "133", "134", "135", "136", "137", "138",
		"139", "147", "150", "151", "152", "153", "155", "156", "157", "158", "159", "186",
		"187", "188"}
	f := prefix[rand.Intn(len(prefix)-1)]
	s := fmt.Sprintf("%08d", rand.Int63n(99999999))
	return f + s
}

func randChatID() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 32
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
