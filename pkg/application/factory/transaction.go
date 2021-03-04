package factory

import (
	"github.com/linmadan/egglib-go/core/application"
	pG "github.com/linmadan/egglib-go/transaction/pg"
	"github.com/linmadan/mmm-newdemo/pkg/infrastructure/pg"
)

func CreateTransactionContext(options map[string]interface{}) (application.TransactionContext, error) {
	return pG.NewPGTransactionContext(pg.DB), nil
}
