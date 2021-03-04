package controllers

import (
	"github.com/linmadan/egglib-go/web/beego"
	"github.com/linmadan/mmm-newdemo/pkg/application/demo/command"
	"github.com/linmadan/mmm-newdemo/pkg/application/demo/query"
	"github.com/linmadan/mmm-newdemo/pkg/application/demo/service"
)

type DemoAdminController struct {
	beego.BaseController
}

func (controller *DemoAdminController) CreateDemo() {
	demoService := service.NewDemoService(nil)
	createDemoCommand := &command.CreateDemoCommand{}
	controller.Unmarshal(createDemoCommand)
	data, err := demoService.CreateDemo(createDemoCommand)
	controller.Response(data, err)
}

func (controller *DemoAdminController) UpdateDemo() {
	demoService := service.NewDemoService(nil)
	updateDemoCommand := &command.UpdateDemoCommand{}
	controller.Unmarshal(updateDemoCommand)
	demoId, _ := controller.GetInt64(":demoId")
	updateDemoCommand.DemoId = demoId
	data, err := demoService.UpdateDemo(updateDemoCommand)
	controller.Response(data, err)
}

func (controller *DemoAdminController) GetDemo() {
	demoService := service.NewDemoService(nil)
	getDemoQuery := &query.GetDemoQuery{}
	demoId, _ := controller.GetInt64(":demoId")
	getDemoQuery.DemoId = demoId
	data, err := demoService.GetDemo(getDemoQuery)
	controller.Response(data, err)
}

func (controller *DemoAdminController) RemoveDemo() {
	demoService := service.NewDemoService(nil)
	removeDemoCommand := &command.RemoveDemoCommand{}
	controller.Unmarshal(removeDemoCommand)
	demoId, _ := controller.GetInt64(":demoId")
	removeDemoCommand.DemoId = demoId
	data, err := demoService.RemoveDemo(removeDemoCommand)
	controller.Response(data, err)
}

func (controller *DemoAdminController) ListDemo() {
	demoService := service.NewDemoService(nil)
	listDemoQuery := &query.ListDemoQuery{}
	offset, _ := controller.GetInt("offset")
	listDemoQuery.Offset = offset
	limit, _ := controller.GetInt("limit")
	listDemoQuery.Limit = limit
	data, err := demoService.ListDemo(listDemoQuery)
	controller.Response(data, err)
}
