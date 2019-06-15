package WechatWork_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "github.com/joho/godotenv/autoload"
)

func TestWechatWorkGo(t *testing.T) {

	RegisterFailHandler(Fail)
	RunSpecs(t, "WechatWorkGo Suite")
}
