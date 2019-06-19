package contact_test

import (
	"fmt"
	"os"

	// wechatwork "github.com/dfang/wechat-work-go"
	wechatwork "github.com/dfang/wechat-work-go"
	"github.com/dfang/wechat-work-go/contact"
	"github.com/dfang/wechat-work-go/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("成员管理 API", func() {
	var c contact.Contact
	testDepartmentID := 9999
	BeforeEach(func() {
		corpID := os.Getenv("CORP_ID")
		client := wechatwork.New(corpID)
		contactAppSecret := os.Getenv("CONTACT_APP_SECRET")
		// 关于创建成员（客服答复）
		// 目前只能使用通讯录的secret 获取token进行创建  其他的secret是没有创建成员的权限的
		// 获取路径：通讯录管理secret。在“管理工具”-“通讯录同步”里面查看（需开启“API接口同步”）
		app := client.WithApp(contactAppSecret, 0)
		c = contact.Contact{
			App: app,
		}
	})

	Describe("创建(C)", func() {

		BeforeEach(func() {
			// 清理先
			clearDepartment(c, testDepartmentID)
		})

		AfterEach(func() {
			// 处理后事
			clearDepartment(c, testDepartmentID)
		})

		Context("不存在的才能创建", func() {
			It("创建部门", func() {
				var data = models.ReqCreateDepartment{
					Name:     "测试部门",
					ParentID: 1,
					Order:    1,
					ID:       testDepartmentID,
				}
				result, _ := c.CreateDepartment(data)
				fmt.Println(result)
				Expect(result.ErrCode).To(Equal(0))
				Expect(result.ID).To(Equal(9999))
			})
		})

		Context("先有部门，才能创建成员", func() {
			BeforeEach(func() {
				createTestDepartment(c, testDepartmentID)
			})

			It("创建成员", func() {
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

				result1, _ := c.CreateMember(u1)
				result2, _ := c.CreateMember(u2)
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

		AfterEach(func() {
			// 处理后事
			clearDepartment(c, testDepartmentID)
		})

		Context("已经存在的部门和成员才能查询、更新", func() {

			It("获取成员", func() {
				result, _ := c.GetMember("zhangsan")
				Expect(result.UserID).To(Equal("zhangsan"))
			})

			It("更新成员", func() {
				u := models.Member{
					UserID:     "zhangsan",
					Name:       "张三三",
					Mobile:     "12345678911",
					Department: []int{testDepartmentID},
				}
				result, _ := c.UpdateMember(u)
				Expect(result.ErrCode).To(Equal(0))
			})

			It("部门列表", func() {
				result, _ := c.ListDepartments(0)
				Expect(result.ErrCode).To(Equal(0))
				Expect(len(result.Department)).To(BeNumerically(">", 0))
				// Ω(len(result.Department)).Should(BeNumerically(">", 0))
			})

			It("更新部门", func() {
				m := models.ReqUpdateDepartment{
					Name:     "测试部门222",
					ParentID: 1,
					Order:    1,
					ID:       testDepartmentID,
				}
				result, _ := c.UpdateDepartment(m)
				Expect(result.ErrCode).To(Equal(0))
			})

			// TODO 需要处理部门不存在的情况
			It("获取部门成员", func() {
				result, _ := c.ListMembers(testDepartmentID, 0)
				// if result.ErrCode != 60003 { }
				fmt.Println(result)
				Expect(len(result.UserList)).To(BeNumerically(">=", 1))
			})

		})

		Context("已经存在的部门和成员删除", func() {
			It("删除部门", func() {
				// 有成员不能删除部门
				// 所以先删完测试用户 再删部门

				// var userIDs []string
				d1, _ := c.ListMembers(testDepartmentID, 0)

				for _, m := range d1.UserList {
					// userIDs = append(userIDs, m.UserID)
					// 批量删除成员接口暂未实现
					c.DeleteMember(m.UserID)
				}

				result, _ := c.DeleteDepartment(testDepartmentID)
				Expect(result.ErrCode).To(Equal(0))
			})

			It("删除成员", func() {
				result, _ := c.DeleteMember("zhangsan")
				Expect(result.ErrCode).To(Equal(0))
			})
		})
	})
})

// clearDepartment 清理部门
//
// a test_helper that clear members in department, then delete the department
func clearDepartment(c contact.Contact, testDepartmentID int) {
	d1, _ := c.ListMembers(testDepartmentID, 0)
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

// createTestDepartment 创建测试部门
//
// a test_helper that create a test department with id: 9999
func createTestDepartment(c contact.Contact, testDepartmentID int) {
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

func createTestUserInDepartment(c contact.Contact, testDepartmentID int) {
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
}
