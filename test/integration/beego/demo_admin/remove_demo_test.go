package demo_admin

import (
	"github.com/go-pg/pg/v10"
	"net/http"

	"github.com/gavv/httpexpect"
	pG "github.com/linmadan/mmm-newdemo/pkg/infrastructure/pg"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("移除例子服务", func() {
	var demoId int64
	BeforeEach(func() {
		_, err := pG.DB.QueryOne(
			pg.Scan(&demoId),
			"INSERT INTO demos (demo_id, demo_name) VALUES (?, ?) RETURNING demo_id",
			1, "testDemoName")
		Expect(err).NotTo(HaveOccurred())
	})
	Describe("根据参数移除例子服务", func() {
		Context("传入有效的demoId", func() {
			It("返回被移除例子的数据", func() {
				httpExpect := httpexpect.New(GinkgoT(), server.URL)
				httpExpect.DELETE("/demos/1").
					Expect().
					Status(http.StatusOK).
					JSON().
					Object().
					ContainsKey("code").ValueEqual("code", 0).
					ContainsKey("msg").ValueEqual("msg", "ok").
					ContainsKey("data").Value("data").Object()
			})
		})
	})
	AfterEach(func() {
		_, err := pG.DB.Exec("DELETE FROM demos WHERE true")
		Expect(err).NotTo(HaveOccurred())
	})
})
