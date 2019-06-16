package wechatwork_test

import (
	"fmt"

	wechatwork "github.com/dfang/wechat-work-go"
	"github.com/dfang/wechat-work-go/models"
	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/resty.v1"
)

var _ = Describe("Access Token", func() {
	BeforeEach(func() {
		// httpmock.ActivateNonDefault(resty.DefaultClient.GetClient())
		httpmock.ActivateNonDefault(wechatwork.NewDefaultRestyClient().GetClient())
	})

	It("should retry when wxqy returns system is busy", func() {
		fixture := `{"errcode": -1, "errmsg": "system is busy"}`
		responder := httpmock.NewStringResponder(200, fixture)
		fakeUrl := "https://api.mybiz.com/articles.json"
		httpmock.RegisterResponder("GET", fakeUrl, responder)
		// app.SyncAccessToken()
		var data models.RespAccessToken
		resp, err := resty.R().SetResult(&data).Get(fakeUrl)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("%+v\n", data)
		fmt.Println(string(resp.Body()))
		fmt.Println(data.AccessToken)
		fmt.Println(data.ExpiresInSecs)
		fmt.Println(data.ErrCode)
		fmt.Println(data.ErrMsg)
		// Expect(resp.Status()).To(Equal(200))
		// TODO: Mock here
		// return -1
	})

	It("should get access token", func() {
		fmt.Println("token", app.AccessToken)
		fmt.Println("expires_in", app.ExpiresIn)

		app.SpawnAccessTokenRefresher()

		// time.Sleep(time.Second * 600)
		// app.SyncAccessToken()
		fmt.Println("###asdfasdf")
		fmt.Println(app.AccessToken)
		Expect(app.AccessToken).NotTo(BeNil())
	})

	It("运行了SyncAccessToken后, app.AccessToken应该不为空", func() {
		fmt.Println("token", app.AccessToken)
		fmt.Println("expires_in", app.ExpiresIn)

		// app.SyncAccessToken()
		app.SpawnAccessTokenRefresher()
		fmt.Println("###asdfasdf")
		fmt.Println(app.AccessToken)
		Expect(app.AccessToken).NotTo(BeNil())
	})

})
