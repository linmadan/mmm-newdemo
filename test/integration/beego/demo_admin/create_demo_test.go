package demo_admin

import (
	"net/http"

	"github.com/gavv/httpexpect"
	pG "github.com/linmadan/mmm-newdemo/pkg/infrastructure/pg"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("创建例子服务", func() {
	Describe("提交数据创建例子服务", func() {
		Context("提交正确的新例子数据", func() {
			It("返回例子数据", func() {
				httpExpect := httpexpect.New(GinkgoT(), server.URL)
				body := map[string]interface{}{
					"demoName": "演示用例",
				}
				httpExpect.POST("/demos/").
					WithJSON(body).
					Expect().
					Status(http.StatusOK).
					JSON().
					Object().
					ContainsKey("code").ValueEqual("code", 0).
					ContainsKey("msg").ValueEqual("msg", "ok").
					ContainsKey("data").Value("data").Object().
					ContainsKey("demoId").ValueNotEqual("demoId", BeZero()).
					ContainsKey("demoName").ValueEqual("demoName", "演示用例")
			})
		})
	})
	AfterEach(func() {
		_, err := pG.DB.Exec("DELETE FROM demos WHERE true")
		Expect(err).NotTo(HaveOccurred())
	})
})
