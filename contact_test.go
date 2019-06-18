package wechatwork_test

import (
	"fmt"
	"os"

	// wechatwork "github.com/dfang/wechat-work-go"
	wechatwork "github.com/dfang/wechat-work-go"
	"github.com/dfang/wechat-work-go/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("成员管理", func() {
	BeforeEach(func() {
		corpID := os.Getenv("CORP_ID")
		client := wechatwork.New(corpID)
		app = client.WithApp("itvtE93S-5QoDunx5yBUhJTccGetSks4Ye6CSs5k5OA", 0)
	})

	Context("基本功能", func() {
		It("创建部门", func() {
			var data = models.ReqCreateDepartment{
				Name:     "测试部门",
				Parentid: 1,
				Order:    1,
				ID:       9999,
			}

			fmt.Println(data)
			result, _ := app.CreateDepartment(data)
			fmt.Println(result)

			Expect(result.ErrCode).To(Equal(0))
			Expect(result.ID).To(Equal(9999))
		})

		It("创建成员", func() {
			var data = models.ReqMemberCreate{
				UserID:     "zhangsan",
				Name:       "张三",
				Department: []int{9999},
				Mobile:     "12345678901",
			}

			fmt.Println(data)
			result, _ := app.CreateMember(data)
			fmt.Println(result)

			Expect(result.ErrCode).To(Equal(0))
		})

		It("获取成员", func() {
			result, _ := app.GetMember("zhangsan")
			fmt.Println(result)
			Expect(result.UserID).To(Equal("zhangsan"))
		})

		It("更新成员", func() {
			m := models.Member{
				UserID:     "zhangsan",
				Name:       "张三三",
				Department: []int{9999},
				Mobile:     "12345678901",
			}
			result, _ := app.UpdateMember(m)
			Expect(result.ErrCode).To(Equal(0))
		})

		It("获取部门成员", func() {
			result, _ := app.ListMembers("9999", false)
			fmt.Println(result)
			Expect(len(result.Userlist)).To(Equal(1))
		})

		It("删除成员", func() {
			// result, _ := app.DeleteMember("zhangsan")
			// Expect(result.ErrCode).To(Equal(0))
		})

		It("更新部门", func() {
			m := models.ReqUpdateDepartment{
				Name:     "测试部门222",
				Parentid: 1,
				Order:    1,
				ID:       9999,
			}
			result, _ := app.UpdateDepartment(m)
			Expect(result.ErrCode).To(Equal(0))
		})

		It("部门列表", func() {
			result, _ := app.ListDepartments("0")
			Expect(result.ErrCode).To(Equal(0))
			Expect(len(result.Department)).To(BeNumerically(">", 0))
			// Ω(len(result.Department)).Should(BeNumerically(">", 0))
		})
	})
})
