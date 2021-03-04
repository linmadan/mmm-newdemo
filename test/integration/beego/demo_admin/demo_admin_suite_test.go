package demo_admin

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/linmadan/mmm-newdemo/pkg/infrastructure/pg"
	_ "github.com/linmadan/mmm-newdemo/pkg/port/beego"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDemoAdmin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Beego Port DemoAdmin Correlations Test Case Suite")
}

var handler http.Handler
var server *httptest.Server

var _ = BeforeSuite(func() {
	handler = web.BeeApp.Handlers
	server = httptest.NewServer(handler)
})

var _ = AfterSuite(func() {
	server.Close()
})
