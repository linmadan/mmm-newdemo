package beego

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/linmadan/egglib-go/web/beego/filters"
	"os"
	"strconv"

	. "github.com/linmadan/mmm-newdemo/pkg/log"
	_ "github.com/linmadan/mmm-newdemo/pkg/port/beego/routers"
)

func init() {
	web.BConfig.AppName = "mmm-newdemo"
	web.BConfig.CopyRequestBody = true
	web.BConfig.RunMode = "dev"
	web.BConfig.Listen.HTTPPort = 8080
	web.BConfig.Listen.EnableAdmin = false
	web.BConfig.WebConfig.CommentRouterPath = "/pkg/port/beego"
	if os.Getenv("RUN_MODE") != "" {
		web.BConfig.RunMode = os.Getenv("RUN_MODE")
	}
	if os.Getenv("HTTP_PORT") != "" {
		portStr := os.Getenv("HTTP_PORT")
		if port, err := strconv.Atoi(portStr); err == nil {
			web.BConfig.Listen.HTTPPort = port
		}
	}
	web.InsertFilter("/*", web.BeforeExec, filters.AllowCors())
	web.InsertFilter("/*", web.BeforeExec, filters.CreateRequstLogFilter(Logger))
	web.InsertFilter("/*", web.AfterExec, filters.CreateResponseLogFilter(Logger), web.WithReturnOnOutput(false))
}
