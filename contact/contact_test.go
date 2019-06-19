package contact_test

import (
	"fmt"

	// wechatwork "github.com/dfang/wechat-work-go"

	"github.com/dfang/wechat-work-go/contact"
	"github.com/dfang/wechat-work-go/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Ginkgo
// You use Describe blocks to describe the individual behaviors of your code and
// Context blocks to exercise those behaviors under different circumstances.

var _ = Describe("成员管理 API", func() {
	var data models.ReqCreateDepartment
	var result models.RespCreateDepartment
	testDepartmentID := 9999

	BeforeEach(func() {
		// 清理先
		clearDepartment(c, testDepartmentID)
	})

	AfterEach(func() {
		clearDepartment(c, testDepartmentID)
	})

	Describe("创建(C)", func() {

		BeforeEach(func() {
			data = models.ReqCreateDepartment{
				Name:     "测试部门",
				ParentID: 1,
				Order:    1,
				ID:       testDepartmentID,
			}
			// fmt.Println("result .....")
			// fmt.Println(result)
		})

		AfterEach(func() {
			// 处理后事
			// clearDepartment(c, testDepartmentID)
		})

		JustBeforeEach(func() {
			result, _ = c.CreateDepartment(data)
			// fmt.Println(data)
			// fmt.Println(result)
			// fmt.Println(c)
		})

		Context("不存在的才能创建", func() {
			It("创建部门", func() {
				Expect(result.ErrCode).To(Equal(0))
				Expect(result.ID).To(Equal(9999))
			})
		})

		Context("先有部门，才能创建成员", func() {
			var result1 models.RespMemberCreate
			var result2 models.RespMemberCreate
			BeforeEach(func() {
				createTestDepartment(c, testDepartmentID)
				var u1 = models.ReqMemberCreate{
					UserID:     "zhangsan",
					Name:       "张三",
					Department: []int{testDepartmentID},
					Mobile:     "12345678901",
				}
				var u2 = models.ReqMemberCreate{
					UserID:     "lisi",
					Name:       "李四",
					Department: []int{testDepartmentID},
					Mobile:     "12345678989",
				}

				result1, _ = c.CreateMember(u1)
				result2, _ = c.CreateMember(u2)
			})

			It("创建成员", func() {
				Expect(result1.ErrCode).To(Equal(0))
				Expect(result2.ErrCode).To(Equal(0))
			})
		})
	})

	Describe("RUD(GET、UPDATE、DELETE）", func() {

		BeforeEach(func() {
			// 创建
			clearDepartment(c, testDepartmentID)
			createTestDepartment(c, testDepartmentID)
			// 创建测试用户，以供查询、更新、删除
			var u1 = models.ReqMemberCreate{
				UserID:     "zhangsan",
				Name:       "张三",
				Department: []int{testDepartmentID},
				Mobile:     "12345678901",
			}
			var u2 = models.ReqMemberCreate{
				UserID:     "lisi",
				Name:       "李四",
				Department: []int{testDepartmentID},
				Mobile:     "12345678989",
			}
			c.CreateMember(u1)
			c.CreateMember(u2)
		})

		// AfterEach(func() {
		// 	// 处理后事
		// 	clearDepartment(c, testDepartmentID)
		// })
		Context("已经存在的部门和成员才能查询、更新", func() {
			var result models.RespMemberGet
			var result2 models.RespCommon
			u := models.Member{
				UserID:     "zhangsan",
				Name:       "张三三",
				Mobile:     "12345678911",
				Department: []int{testDepartmentID},
			}

			BeforeEach(func() {
				result, _ = c.GetMember("zhangsan")
				result2, _ = c.UpdateMember(u)
			})

			It("获取成员", func() {
				Expect(result.UserID).To(Equal("zhangsan"))
			})

			It("更新成员", func() {

				Expect(result2.ErrCode).To(Equal(0))
			})
		})

		Context("已经存在的部门和成员才能查询、更新", func() {
			var result models.RespListDepartments
			BeforeEach(func() {
				result, _ = c.ListDepartments(0)
			})
			It("部门列表", func() {
				Expect(result.ErrCode).To(Equal(0))
				Expect(len(result.Department)).To(BeNumerically(">", 0))
				// Ω(len(result.Department)).Should(BeNumerically(">", 0))
			})
		})

		Context("已经存在的部门更新", func() {
			m := models.ReqUpdateDepartment{
				Name:     "测试部门222",
				ParentID: 1,
				Order:    1,
				ID:       testDepartmentID,
			}
			var result models.RespCommon
			BeforeEach(func() {
				result, _ = c.UpdateDepartment(m)
			})
			It("更新部门", func() {
				Expect(result.ErrCode).To(Equal(0))
			})
		})

		Context("已经存在的部门才能获取", func() {
			var result models.RespListMembers
			BeforeEach(func() {
				result, _ = c.ListMembers(testDepartmentID, 0)
			})
			// TODO 需要处理部门不存在的情况
			It("获取部门成员", func() {
				// if result.ErrCode != 60003 { }
				// fmt.Println(result)
				Expect(len(result.UserList)).To(BeNumerically(">=", 1))
			})

		})

		Context("已经存在的成员才能删除", func() {
			var result models.RespCommon
			BeforeEach(func() {
				result, _ = c.DeleteMember("zhangsan")
			})
			It("删除成员", func() {
				Expect(result.ErrCode).To(Equal(0))
			})
		})

		// Context("已经存在的部门才能删除", func() {
		// 	var result models.RespListMembers
		// 	BeforeEach(func() {
		// 		result, _ = c.ListMembers(testDepartmentID, 0)
		// 	})

		// 	It("删除部门", func() {
		// 		// 有成员不能删除部门
		// 		// 所以先删完测试用户 再删部门
		// 		// var userIDs []string
		// 		d1, _ := c.ListMembers(testDepartmentID, 0)

		// 		for _, m := range d1.UserList {
		// 			// userIDs = append(userIDs, m.UserID)
		// 			// 批量删除成员接口暂未实现
		// 			c.DeleteMember(m.UserID)
		// 		}

		// 		result, _ := c.DeleteDepartment(testDepartmentID)
		// 		Expect(result.ErrCode).To(Equal(0))
		// 	})
		// })

	})
})

// clearDepartment 清理部门
//
// a test_helper that clear members in department, then delete the department
func clearDepartment(c *contact.Contact, testDepartmentID int) {
	d1, _ := c.ListMembers(testDepartmentID, 0)
	fmt.Println("%%%%%%%%%%%")
	fmt.Println(d1)
	// 60003 部门不存在
	if d1.ErrCode != 60003 {
		var ulist []string
		for _, m := range d1.UserList {
			ulist = append(ulist, m.UserID)
		}
		if len(ulist) > 0 {
			req := models.ReqBatchDeleteMembers{
				UserIDList: ulist,
			}
			c.DeleteMembers(req)
		}
		result2, _ := c.DeleteDepartment(testDepartmentID)
		if result2.ErrCode == 0 {
			fmt.Println("BeforeEach Prepared")
		}
	}
}

// createTestDepartment 创建测试部门
//
// a test_helper that create a test department with id: 9999
func createTestDepartment(c *contact.Contact, testDepartmentID int) {
	var data = models.ReqCreateDepartment{
		Name:     "测试部门",
		ParentID: 1,
		Order:    1,
		ID:       testDepartmentID,
	}
	result, _ := c.CreateDepartment(data)
	if result.ErrCode == 0 {
		fmt.Println("BeforeEach Prepared")
	}
}

// func createTestUserInDepartment(c *contact.Contact, testDepartmentID int) {
// 	var u1 = models.ReqMemberCreate{
// 		UserID:     "zhangsan",
// 		Name:       "张三",
// 		Department: []int{testDepartmentID},
// 		Mobile:     "12345678901",
// 	}

// 	var u2 = models.ReqMemberCreate{
// 		UserID:     "lisi",
// 		Name:       "李四",
// 		Department: []int{testDepartmentID},
// 		Mobile:     "12345678989",
// 	}

// 	c.CreateMember(u1)
// 	c.CreateMember(u2)
// }
