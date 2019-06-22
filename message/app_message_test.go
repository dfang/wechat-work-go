package message_test

import (
	"github.com/dfang/wechat-work-go/message"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AppMessage", func() {
	Context("发送应用信息", func() {
		var result message.RespSendAppMsg
		JustBeforeEach(func() {
			msg1 := message.NewAppTextMessage(msg.App.AgentID, "this a test msage", 0)
			msg1.ToUser = "@all"

			result, _ = msg.SendAppMessage(msg1)
		})

		It("发送应用信息", func() {
			Expect(result.ErrCode).To(Equal(0))
		})
	})
})
