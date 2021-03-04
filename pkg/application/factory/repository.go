package factory

import (
	"github.com/linmadan/egglib-go/transaction/pg"
	"github.com/linmadan/mmm-newdemo/pkg/domain"
	"github.com/linmadan/mmm-newdemo/pkg/infrastructure/repository"
)

func CreateDemoRepository(options map[string]interface{}) (domain.DemoRepository, error) {
	var transactionContext *pg.TransactionContext
	if value, ok := options["transactionContext"]; ok {
		transactionContext = value.(*pg.TransactionContext)
	}
	return repository.NewDemoRepository(transactionContext)
}
