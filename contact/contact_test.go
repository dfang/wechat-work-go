package contact_test

import (

	// wechatwork "github.com/dfang/wechat-work-go"

	"github.com/dfang/wechat-work-go/contact"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Ginkgo
// You use Describe blocks to describe the individual behaviors of your code and
// Context blocks to exercise those behaviors under different circumstances.

var _ = Describe("成员管理 API", func() {
	// var data contact.ReqCreateDepartment
	// var result contact.RespCreateDepartment

	testDepartmentID := 99999
	var resp1 contact.RespCommon
	JustBeforeEach(func() {
		createTestDepartment(c, testDepartmentID)
		createTestUsersInDepartment(c, testDepartmentID)

		body := contact.ReqUpdateDepartment{
			ID:       testDepartmentID,
			ParentID: 1,
			Name:     "X 部门",
		}
		resp1, _ = c.UpdateDepartment(body)
	})

	Context("创建", func() {

		It("能更新部门信息", func() {
			Expect(resp1.ErrCode).To(Equal(0))
			By("能更新部门信息，说明部门存在，创建不用测试")
		})

		Context("已经存在的部门和成员才能查询、更新", func() {
			var result contact.RespMemberGet
			var result2 contact.RespCommon
			u := contact.Member{
				UserID: "zhangsan",
				Name:   "张三三",
				Mobile: "12345678911",
				// Department: []int{*testDepartmentID},
			}
			JustBeforeEach(func() {
				result, _ = c.GetMember("zhangsan")
				result2, _ = c.UpdateMember(u)
			})

			It("获取成员", func() {
				Expect(result.UserID).To(Equal("zhangsan"))
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
func createTestUsersInDepartment(c *contact.Contact, testDepartmentID int) {
	var u1 = contact.ReqMemberCreate{
		UserID:     "zhangsan",
		Name:       "张三",
		Department: []int{testDepartmentID},
		Mobile:     "12345678901",
		Enable:     1,
	}

	var u2 = contact.ReqMemberCreate{
		UserID:     "lisi",
		Name:       "李四",
		Department: []int{testDepartmentID},
		Mobile:     "12345678989",
		Enable:     1,
	}

	var u3 = contact.ReqMemberCreate{
		UserID:         "df1228",
		Name:           "df1228",
		Alias:          "我是部门老大 哈哈哈哈",
		Position:       "部门BOSS",
		Department:     []int{testDepartmentID},
		Mobile:         "15618903181",
		Gender:         "男",
		Enable:         1,
		IsLeaderInDept: []int{1},
	}

	_, _ = c.CreateMember(u1)
	_, _ = c.CreateMember(u2)
	_, _ = c.CreateMember(u3)

	By("Test Users Created")
}
