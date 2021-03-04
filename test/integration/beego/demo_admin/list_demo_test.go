package demo_admin

import (
	"github.com/go-pg/pg/v10"
	"net/http"

	"github.com/gavv/httpexpect"
	pG "github.com/linmadan/mmm-newdemo/pkg/infrastructure/pg"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("返回例子服务列表", func() {
	var demoId int64
	BeforeEach(func() {
		_, err := pG.DB.QueryOne(
			pg.Scan(&demoId),
			"INSERT INTO demos (demo_id, demo_name) VALUES (?, ?) RETURNING demo_id",
			1, "testDemoName")
		Expect(err).NotTo(HaveOccurred())
	})
	Describe("根据参数返回例子列表", func() {
		Context("传入有效的参数", func() {
			It("返回例子数据列表", func() {
				httpExpect := httpexpect.New(GinkgoT(), server.URL)
				httpExpect.GET("/demos/").
					WithQuery("offset", 0).
					WithQuery("limit", 20).
					Expect().
					Status(http.StatusOK).
					JSON().
					Object().
					ContainsKey("code").ValueEqual("code", 0).
					ContainsKey("msg").ValueEqual("msg", "ok").
					ContainsKey("data").Value("data").Object().
					ContainsKey("count").ValueEqual("count", 1).
					ContainsKey("demos").Value("demos").Array()
			})
		})
	})
	AfterEach(func() {
		_, err := pG.DB.Exec("DELETE FROM demos WHERE true")
		Expect(err).NotTo(HaveOccurred())
	})
})
