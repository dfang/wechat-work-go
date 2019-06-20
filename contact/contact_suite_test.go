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

// var testDepartmentID *int
var testDepartmentData *contact.ReqCreateDepartment

func TestContact(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contact Suite")
}

var _ = BeforeSuite(func() {
	corpID := os.Getenv("CORP_ID")
	contactAppSecret := os.Getenv("CONTACT_APP_SECRET")
	agentID, _ := strconv.ParseInt(os.Getenv("AGENT_ID"), 10, 64)

	// testDepartmentID = 99999

	client := wechatwork.New(corpID)
	app := client.NewApp(contactAppSecret, agentID)
	c = contact.WithApp(app)

	// Clear test department
	clearDepartment(c, 99999)

	testDepartmentData = &contact.ReqCreateDepartment{
		Name:     "测试部门",
		ParentID: 1,
		Order:    1,
		ID:       99999,
	}
})

var _ = AfterSuite(func() {
	clearDepartment(c, 99999)
})
