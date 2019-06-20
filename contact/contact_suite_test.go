package contact_test

import (
	"os"
	"strconv"
	"testing"

	wechatwork "github.com/dfang/wechat-work-go"
	"github.com/dfang/wechat-work-go/contact"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "github.com/joho/godotenv/autoload"
)

// var app *wechatwork.App
var c *contact.Contact

func TestContact(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contact Suite")
}

var _ = BeforeSuite(func() {
	corpID := os.Getenv("CORP_ID")
	contactAppSecret := os.Getenv("CONTACT_APP_SECRET")
	agentID, _ := strconv.ParseInt(os.Getenv("AGENT_ID"), 10, 64)

	client := wechatwork.New(corpID)
	app := client.NewApp(contactAppSecret, agentID)
	c = &contact.Contact{
		App: app,
	}
})
