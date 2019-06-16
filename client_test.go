package wechatwork_test

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/dfang/wechat-work-go/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/resty.v1"
)

var _ = Describe("Client", func() {
	var req *resty.Request
	var u *url.URL
	var err error
	BeforeEach(func() {
		apiPath := "/cgi-bin/agent/get"
		qs := models.ReqAgentGet{
			AgentID: strconv.FormatInt(app.AgentID, 10),
		}
		req = app.NewRequest(apiPath, qs, true)
		fmt.Println("#####")
		fmt.Println(req)
		fmt.Println(req.URL)
		fmt.Println(req.QueryParam.Get("access_token"))

		u, err = url.Parse(req.URL)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(GinkgoWriter, "Some log text: %s\n", "ad")
	})

	Context("should return resty.Request", func() {

		It("should has access_token in querystring and should not be empty string", func() {
			Expect(u.Query().Get("access_token")).NotTo(BeNil())
			Expect(u.Query().Get("access_token")).NotTo(Equal(""))
		})
	})

})

// var _ = Describe("Articles", func() {
// 	It("returns a list of articles", func() {
// 		fixture := `{"status":{"message": "Your message", "code": 200}}`
// 		responder := httpmock.NewStringResponder(200, fixture)
// 		fakeUrl := "https://api.mybiz.com/articles.json"
// 		httpmock.RegisterResponder("GET", fakeUrl, responder)

// 		// fetch the article into struct
// 		articleObject := &models.Article{}
// 		_, err := resty.R().SetResult(articleObject).Get(fakeUrl)

// 		// do stuff with the article object ...
// 	})
// })
