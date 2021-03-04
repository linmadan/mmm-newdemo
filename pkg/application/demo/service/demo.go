package service

import (
	"fmt"
	"github.com/linmadan/egglib-go/core/application"
	"github.com/linmadan/egglib-go/utils/tool_funs"
	"github.com/linmadan/mmm-newdemo/pkg/application/demo/command"
	"github.com/linmadan/mmm-newdemo/pkg/application/demo/query"
	"github.com/linmadan/mmm-newdemo/pkg/application/factory"
	"github.com/linmadan/mmm-newdemo/pkg/domain"
)

// 例子服务
type DemoService struct {
}

// 创建例子服务
func (demoService *DemoService) CreateDemo(createDemoCommand *command.CreateDemoCommand) (interface{}, error) {
	if err := createDemoCommand.ValidateCommand(); err != nil {
		return nil, application.ThrowError(application.ARG_ERROR, err.Error())
	}
	transactionContext, err := factory.CreateTransactionContext(nil)
	if err != nil {
		return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
	}
	if err := transactionContext.StartTransaction(); err != nil {
		return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
	}
	defer func() {
		transactionContext.RollbackTransaction()
	}()
	newDemo := &domain.Demo{
		DemoName: createDemoCommand.DemoName,
	}
	var demoRepository domain.DemoRepository
	if value, err := factory.CreateDemoRepository(map[string]interface{}{
		"transactionContext": transactionContext,
	}); err != nil {
		return nil, application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
	} else {
		demoRepository = value
	}
	if demo, err := demoRepository.Save(newDemo); err != nil {
		return nil, application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
	} else {
		if err := transactionContext.CommitTransaction(); err != nil {
			return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
		}
		return demo, nil
	}
}

// 返回例子服务
func (demoService *DemoService) GetDemo(getDemoQuery *query.GetDemoQuery) (interface{}, error) {
	if err := getDemoQuery.ValidateQuery(); err != nil {
		return nil, application.ThrowError(application.ARG_ERROR, err.Error())
	}
	transactionContext, err := factory.CreateTransactionContext(nil)
	if err != nil {
		return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
	}
	if err := transactionContext.StartTransaction(); err != nil {
		return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
	}
	defer func() {
		transactionContext.RollbackTransaction()
	}()
	var demoRepository domain.DemoRepository
	if value, err := factory.CreateDemoRepository(map[string]interface{}{
		"transactionContext": transactionContext,
	}); err != nil {
		return nil, application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
	} else {
		demoRepository = value
	}
	demo, err := demoRepository.FindOne(map[string]interface{}{"demoId": getDemoQuery.DemoId})
	if err != nil {
		return nil, application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
	}
	if demo == nil {
		return nil, application.ThrowError(application.RES_NO_FIND_ERROR, fmt.Sprintf("%s", string(getDemoQuery.DemoId)))
	} else {
		if err := transactionContext.CommitTransaction(); err != nil {
			return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
		}
		return demo, nil
	}
}

// 返回例子服务列表
func (demoService *DemoService) ListDemo(listDemoQuery *query.ListDemoQuery) (interface{}, error) {
	if err := listDemoQuery.ValidateQuery(); err != nil {
		return nil, application.ThrowError(application.ARG_ERROR, err.Error())
	}
	transactionContext, err := factory.CreateTransactionContext(nil)
	if err != nil {
		return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
	}
	if err := transactionContext.StartTransaction(); err != nil {
		return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
	}
	defer func() {
		transactionContext.RollbackTransaction()
	}()
	var demoRepository domain.DemoRepository
	if value, err := factory.CreateDemoRepository(map[string]interface{}{
		"transactionContext": transactionContext,
	}); err != nil {
		return nil, application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
	} else {
		demoRepository = value
	}
	if count, demos, err := demoRepository.Find(tool_funs.SimpleStructToMap(listDemoQuery)); err != nil {
		return nil, application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
	} else {
		if err := transactionContext.CommitTransaction(); err != nil {
			return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
		}
		return map[string]interface{}{
			"count": count,
			"demos": demos,
		}, nil
	}
}

// 移除例子服务
func (demoService *DemoService) RemoveDemo(removeDemoCommand *command.RemoveDemoCommand) (interface{}, error) {
	if err := removeDemoCommand.ValidateCommand(); err != nil {
		return nil, application.ThrowError(application.ARG_ERROR, err.Error())
	}
	transactionContext, err := factory.CreateTransactionContext(nil)
	if err != nil {
		return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
	}
	if err := transactionContext.StartTransaction(); err != nil {
		return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
	}
	defer func() {
		transactionContext.RollbackTransaction()
	}()
	var demoRepository domain.DemoRepository
	if value, err := factory.CreateDemoRepository(map[string]interface{}{
		"transactionContext": transactionContext,
	}); err != nil {
		return nil, application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
	} else {
		demoRepository = value
	}
	demo, err := demoRepository.FindOne(map[string]interface{}{"demoId": removeDemoCommand.DemoId})
	if err != nil {
		return nil, application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
	}
	if demo == nil {
		return nil, application.ThrowError(application.RES_NO_FIND_ERROR, fmt.Sprintf("%s", string(removeDemoCommand.DemoId)))
	}
	if demo, err := demoRepository.Remove(demo); err != nil {
		return nil, application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
	} else {
		if err := transactionContext.CommitTransaction(); err != nil {
			return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
		}
		return demo, nil
	}
}

// 更新例子服务
func (demoService *DemoService) UpdateDemo(updateDemoCommand *command.UpdateDemoCommand) (interface{}, error) {
	if err := updateDemoCommand.ValidateCommand(); err != nil {
		return nil, application.ThrowError(application.ARG_ERROR, err.Error())
	}
	transactionContext, err := factory.CreateTransactionContext(nil)
	if err != nil {
		return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
	}
	if err := transactionContext.StartTransaction(); err != nil {
		return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
	}
	defer func() {
		transactionContext.RollbackTransaction()
	}()
	var demoRepository domain.DemoRepository
	if value, err := factory.CreateDemoRepository(map[string]interface{}{
		"transactionContext": transactionContext,
	}); err != nil {
		return nil, application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
	} else {
		demoRepository = value
	}
	demo, err := demoRepository.FindOne(map[string]interface{}{"demoId": updateDemoCommand.DemoId})
	if err != nil {
		return nil, application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
	}
	if demo == nil {
		return nil, application.ThrowError(application.RES_NO_FIND_ERROR, fmt.Sprintf("%s", string(updateDemoCommand.DemoId)))
	}
	if err := demo.Update(tool_funs.SimpleStructToMap(updateDemoCommand)); err != nil {
		return nil, application.ThrowError(application.BUSINESS_ERROR, err.Error())
	}
	if demo, err := demoRepository.Save(demo); err != nil {
		return nil, application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
	} else {
		if err := transactionContext.CommitTransaction(); err != nil {
			return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
		}
		return demo, nil
	}
}

func NewDemoService(options map[string]interface{}) *DemoService {
	newDemoService := &DemoService{}
	return newDemoService
}
