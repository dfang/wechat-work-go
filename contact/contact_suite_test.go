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

var c *contact.Contact
var testDepartmentID *int
var testUserID *string
var testDepartmentData *contact.ReqCreateDepartment

func TestContact(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contact Suite")
}

var _ = BeforeSuite(func() {
	By("BeforeSuite started")

	corpID := os.Getenv("CORP_ID")
	contactAppSecret := os.Getenv("CONTACT_APP_SECRET")
	agentID, _ := strconv.ParseInt(os.Getenv("AGENT_ID"), 10, 64)

	testDepartmentID = intPtr(99999)
	// testDepartmentID = new(int)
	// *testDepartmentID = 99999

	corp := wechatwork.New(corpID)
	app := corp.NewApp(contactAppSecret, agentID)
	c = contact.WithApp(app)

	// Clear test department
	clearDepartment(c, *testDepartmentID)

	testDepartmentData = &contact.ReqCreateDepartment{
		Name:     "测试部门",
		ParentID: 1,
		Order:    1,
		ID:       *testDepartmentID,
	}

	createTestDepartment(c, *testDepartmentID)

	userIDs := createTestUsersInDepartment(c, *testDepartmentID)
	testUserID = strPtr(userIDs[0])

	By("BeforeSuite completed")
})

var _ = AfterSuite(func() {
	By("AfterSuite started")

	// clearDepartment(c, 99999)

	By("AfterSuite started")
})

func intPtr(i int) *int {
	return &i
}

func strPtr(s string) *string {
	return &s
}
