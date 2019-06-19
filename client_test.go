package wechatwork_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Client", func() {
	// var req *resty.Request
	// var u *url.URL
	// var err error
	// var resp *resty.Response
	BeforeEach(func() {
		// apiPath := "/cgi-bin/agent/get"
		// qs := models.ReqAgentGet{
		// 	AgentID: strconv.FormatInt(app.AgentID, 10),
		// }
		// req = app.NewRequest(apiPath, qs, true)
		// fmt.Println(req.URL)
		// u, err = url.Parse(req.URL)
		// if err != nil {
		// 	panic(err)
		// }
		// resp, err = req.Get(req.URL)
		// if err != nil {
		// 	panic(err)
		// }

	})

	Context("should return resty.Request", func() {
		// It("should has correct HostURL", func() {
		// 	Expect(resp.Header().Get("HOST")).To(Equal("https://qyapi.weixin.qq.com"))
		// })

		// It("should has access_token in querystring and should not be empty string", func() {
		// 	Expect(u.Query().Get("access_token")).NotTo(BeNil())
		// 	Expect(u.Query().Get("access_token")).NotTo(Equal(""))
		// })

	})
})
