package WechatWork_test

import (
	"testing"

	WechatWork "github.com/dfang/wechat-work-go"
	c "github.com/smartystreets/goconvey/convey"
)

func TestWechatWorkCtor(t *testing.T) {
	c.Convey("不带参数构造 WechatWork 实例", t, func() {
		a := WechatWork.New("testcorpid")

		c.Convey("corpid 应该正确设置了", func() {
			c.So(a.CorpID, c.ShouldEqual, "testcorpid")
		})
	})
}
