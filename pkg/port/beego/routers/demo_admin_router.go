package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/linmadan/mmm-newdemo/pkg/port/beego/controllers"
)

func init() {
	web.Router("/demos/", &controllers.DemoAdminController{}, "Post:CreateDemo")
	web.Router("/demos/:demoId", &controllers.DemoAdminController{}, "Put:UpdateDemo")
	web.Router("/demos/:demoId", &controllers.DemoAdminController{}, "Get:GetDemo")
	web.Router("/demos/:demoId", &controllers.DemoAdminController{}, "Delete:RemoveDemo")
	web.Router("/demos/", &controllers.DemoAdminController{}, "Get:ListDemo")
}
