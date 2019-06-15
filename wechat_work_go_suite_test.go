package WechatWork_test

import (
	"os"
	"strconv"
	"testing"

	WechatWork "github.com/dfang/wechat-work-go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "github.com/joho/godotenv/autoload"
)

var app *WechatWork.App

func TestWechatWorkGo(t *testing.T) {

	RegisterFailHandler(Fail)
	RunSpecs(t, "WechatWorkGo Suite")
}

var _ = BeforeSuite(func() {
	corpID := os.Getenv("CORP_ID")
	corpSecret := os.Getenv("CORP_SECRET")
	agentID, _ := strconv.ParseInt(os.Getenv("AGENT_ID"), 10, 64)

	client := WechatWork.New(corpID)
	app = client.WithApp(corpSecret, agentID)
})
