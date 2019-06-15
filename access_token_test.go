package WechatWork_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Access Token", func() {

	It("should get access token", func() {
		app.SpawnAccessTokenRefresher()
		Expect(app.AccessToken).NotTo(BeNil())
	})

	It("运行了SyncAccessToken后, app.AccessToken应该不为空", func() {
		app.SyncAccessToken()
		Expect(app.AccessToken).NotTo(BeNil())
	})

})
