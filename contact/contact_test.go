package contact_test

import (

	// wechatwork "github.com/dfang/wechat-work-go"

	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/dfang/wechat-work-go/contact"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Ginkgo
// You use Describe blocks to describe the individual behaviors of your code and
// Context blocks to exercise those behaviors under different circumstances.

var _ = Describe("成员管理 API", func() {
	var resp1 contact.RespCommon

	Context("能更新部门信息", func() {
		JustBeforeEach(func() {
			body := contact.ReqUpdateDepartment{
				ID:       *testDepartmentID,
				ParentID: 1,
				Name:     fmt.Sprintf("X 部门 %d", rand.Intn(10)),
			}
			resp1, _ = c.UpdateDepartment(body)
		})

		It("能更新部门信息", func() {
			Expect(resp1.ErrCode).To(Equal(0))
			By("能更新部门信息，说明部门存在，创建不用测试")
		})
	})

	Context("已经存在的成员才能查询、更新", func() {
		var result contact.RespGetMember
		var result2 contact.RespCommon
		JustBeforeEach(func() {
			u := contact.Member{
				UserID:     *testUserID,
				Name:       randCNName(),
				Mobile:     randCNPhone(),
				Enable:     1,
				Department: []int{*testDepartmentID},
			}
			result, _ = c.GetMember(*testUserID)
			result2, _ = c.UpdateMember(u)
		})

		It("获取成员", func() {
			Expect(result.UserID).To(Equal(*testUserID))
			By("能获取成员信息，说明成员存在，创建不用测试")
		})

		It("更新成员", func() {
			Expect(result2.ErrCode).To(Equal(0))
			By("能更新成员信息，说明成员存在，创建不用测试")
		})
	})

	// 创建和删除部门、成员都不需测试, 因为BeforeSuite、AfterSuite
	// TODO 添加其他api的测试
	// UserIDToOpenID
	// OpenIDToUserID
	// TwoFactorAuth
	// InviteMember

})

// clearDepartment 清理部门
// a test_helper that clear members in department, then delete the department
func clearDepartment(c *contact.Contact, testDepartmentID int) {
	d1, _ := c.ListMembers(testDepartmentID, 0)
	// 60003 部门不存在
	if d1.ErrCode != 60003 {
		var ulist []string
		for _, m := range d1.UserList {
			ulist = append(ulist, m.UserID)
		}
		if len(ulist) > 0 {
			req := contact.ReqBatchDeleteMembers{
				UserIDList: ulist,
			}
			_, _ = c.DeleteMembers(req)
		}
		result2, _ := c.DeleteDepartment(testDepartmentID)
		if result2.ErrCode == 0 {
			By("Department Cleared")
		}
	}
}

// createTestDepartment 创建测试部门
// a test_helper that create a test department with id: 99999
func createTestDepartment(c *contact.Contact, testDepartmentID int) {
	var data = contact.ReqCreateDepartment{
		Name:     "测试部门",
		ParentID: 1,
		Order:    1,
		ID:       testDepartmentID,
	}
	result, _ := c.CreateDepartment(data)
	if result.ErrCode == 0 {
		By("Department Created")
	}
}

// createTestUsersInDepartment  在测试部门里创建一些测试成员
func createTestUsersInDepartment(c *contact.Contact, testDepartmentID int) []string {

	var a []string
	u1, err := createTestUserInDepartment(c, testDepartmentID)
	if err != nil {
		panic("u1")
	}
	a = append(a, u1)
	u2, err := createTestUserInDepartment(c, testDepartmentID)
	if err != nil {
		panic("u2")
	}
	a = append(a, u2)

	u3, err := createTestUserInDepartment(c, testDepartmentID)
	if err != nil {
		panic("u3")
	}
	a = append(a, u3)

	// var u3 = contact.ReqCreateMember{
	// 	UserID:         "df1228",
	// 	Name:           "df1228",
	// 	Alias:          "我是部门老大 哈哈哈哈",
	// 	Position:       "部门BOSS",
	// 	Department:     []int{testDepartmentID},
	// 	Mobile:         "15618903181",
	// 	Gender:         "男",
	// 	Enable:         1,
	// 	IsLeaderInDept: []int{1},
	// }

	By("Test Users Created")

	return a
}

func createTestUserInDepartment(c *contact.Contact, testDepartmentID int) (string, error) {
	var u1 = contact.ReqCreateMember{
		UserID:     randomUserID(),
		Name:       randCNName(),
		Department: []int{testDepartmentID},
		Mobile:     randCNPhone(),
		Enable:     1,
	}
	_, err := c.CreateMember(u1)
	if err != nil {
		return "", err
	}
	return u1.UserID, nil
}

func randomUserID() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

/*
func randomMobilePhone() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("0123456789")
	length := 11
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
*/

func randCNName() string {
	rand.Seed(time.Now().UnixNano())
	var fList []string = []string{"赵", "钱", "孙", "李", "周", "吴", "郑", "王", "冯", "陈", "褚", "卫", "蒋", "沈", "韩", "杨", "朱"}
	var lList []string = []string{"一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}
	f := fList[rand.Intn(len(fList)-1)]
	l := lList[rand.Intn(len(lList)-1)]
	return f + l
}

func randCNPhone() string {
	rand.Seed(time.Now().UnixNano())
	prefix := []string{"130", "131", "132", "133", "134", "135", "136", "137", "138",
		"139", "147", "150", "151", "152", "153", "155", "156", "157", "158", "159", "186",
		"187", "188"}
	f := prefix[rand.Intn(len(prefix)-1)]
	s := fmt.Sprintf("%08d", rand.Int63n(99999999))
	return f + s
}
