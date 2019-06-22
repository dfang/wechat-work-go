package wechatwork_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Access Token", func() {
	BeforeEach(func() {
		// httpmock.ActivateNonDefault(resty.DefaultClient.GetClient())
	})

	Context("Access Token", func() {

		It("should", func() {
			tokn := app.GetAccessToken()
			Expect(tokn).NotTo(BeEmpty())
		})

		// It("shoule panic when either corpid or corpsecret not correct, and both not correct", func() {
		// 	// app.SpawnAccessTokenRefresher()
		// 	// TODO: don't know how to test yet
		// 	// JUST pass a wrong corpid or corpsecret, or both wrong ones
		// 	// and try to request access_token
		// 	// for now, there is a bug for wxqy, if both corpid and corpsecret are not correct
		// 	// getToken returns nothing
		// 	Expect(true).To(Equal(true))
		// })

		// It("should retry when wxqy server returns system is busy", func() {
		// 	// TODO: Mock here
		// })

		// It("should get access token, 运行了SpawnAccessTokenRefresher后, app.AccessToken应该不为空", func() {
		// 	Expect(app.AccessToken).To(Equal(""))
		// 	Expect(app.ExpiresIn).To(Equal(0))

		// 	app.SyncAccessToken()
		// 	// app.SpawnAccessTokenRefresher()
		// 	// time.Sleep(time.Second * 5)

		// 	Expect(app.AccessToken).NotTo(BeNil())
		// 	Expect(app.AccessToken).NotTo(Equal(""))
		// })

	})

})
