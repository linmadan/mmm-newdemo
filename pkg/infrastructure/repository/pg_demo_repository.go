package repository

import (
	"fmt"
	"github.com/go-pg/pg/v10"

	"github.com/linmadan/egglib-go/persistent/pg/sqlbuilder"
	pgTransaction "github.com/linmadan/egglib-go/transaction/pg"
	"github.com/linmadan/egglib-go/utils/snowflake"
	"github.com/linmadan/mmm-newdemo/pkg/domain"
	"github.com/linmadan/mmm-newdemo/pkg/infrastructure/pg/models"
	"github.com/linmadan/mmm-newdemo/pkg/infrastructure/pg/transform"
)

type DemoRepository struct {
	transactionContext *pgTransaction.TransactionContext
}

func (repository *DemoRepository) nextIdentify() (int64, error) {
	IdWorker, err := snowflake.NewIdWorker(1)
	if err != nil {
		return 0, err
	}
	id, err := IdWorker.NextId()
	return id, err
}
func (repository *DemoRepository) Save(demo *domain.Demo) (*domain.Demo, error) {
	sqlBuildFields := []string{
		"demo_id",
		"demo_name",
	}
	insertFieldsSnippet := sqlbuilder.SqlFieldsSnippet(sqlBuildFields)
	insertPlaceHoldersSnippet := sqlbuilder.SqlPlaceHoldersSnippet(sqlBuildFields)
	returningFieldsSnippet := sqlbuilder.SqlFieldsSnippet(sqlBuildFields)
	updateFields := sqlbuilder.RemoveSqlFields(sqlBuildFields, "demo_id")
	updateFieldsSnippet := sqlbuilder.SqlUpdateFieldsSnippet(updateFields)
	tx := repository.transactionContext.PgTx
	if demo.Identify() == nil {
		demoId, err := repository.nextIdentify()
		if err != nil {
			return demo, err
		} else {
			demo.DemoId = demoId
		}
		if _, err := tx.QueryOne(
			pg.Scan(
				&demo.DemoId,
				&demo.DemoName,
			),
			fmt.Sprintf("INSERT INTO demos (%s) VALUES (%s) RETURNING %s", insertFieldsSnippet, insertPlaceHoldersSnippet, returningFieldsSnippet),
			demo.DemoId,
			demo.DemoName,
		); err != nil {
			return demo, err
		}
	} else {
		if _, err := tx.QueryOne(
			pg.Scan(
				&demo.DemoId,
				&demo.DemoName,
			),
			fmt.Sprintf("UPDATE demos SET %s WHERE demo_id=? RETURNING %s", updateFieldsSnippet, returningFieldsSnippet),
			demo.DemoName,
			demo.Identify(),
		); err != nil {
			return demo, err
		}
	}
	return demo, nil
}
func (repository *DemoRepository) Remove(demo *domain.Demo) (*domain.Demo, error) {
	tx := repository.transactionContext.PgTx
	demoModel := new(models.Demo)
	demoModel.DemoId = demo.Identify().(int64)
	if _, err := tx.Model(demoModel).WherePK().Delete(); err != nil {
		return demo, err
	}
	return demo, nil
}
func (repository *DemoRepository) FindOne(queryOptions map[string]interface{}) (*domain.Demo, error) {
	tx := repository.transactionContext.PgTx
	demoModel := new(models.Demo)
	query := sqlbuilder.BuildQuery(tx.Model(demoModel), queryOptions)
	query.SetWhereByQueryOption("demo.demo_id = ?", "demoId")
	if err := query.First(); err != nil {
		if err.Error() == "pg: no rows in result set" {
			return nil, fmt.Errorf("没有此资源")
		} else {
			return nil, err
		}
	}
	if demoModel.DemoId == 0 {
		return nil, nil
	} else {
		return transform.TransformToDemoDomainModelFromPgModels(demoModel)
	}
}
func (repository *DemoRepository) Find(queryOptions map[string]interface{}) (int64, []*domain.Demo, error) {
	tx := repository.transactionContext.PgTx
	var demoModels []*models.Demo
	demos := make([]*domain.Demo, 0)
	query := sqlbuilder.BuildQuery(tx.Model(&demoModels), queryOptions)
	query.SetOffsetAndLimit(20)
	query.SetOrderDirect("demo_id", "DESC")
	if count, err := query.SelectAndCount(); err != nil {
		return 0, demos, err
	} else {
		for _, demoModel := range demoModels {
			if demo, err := transform.TransformToDemoDomainModelFromPgModels(demoModel); err != nil {
				return 0, demos, err
			} else {
				demos = append(demos, demo)
			}
		}
		return int64(count), demos, nil
	}
}
func NewDemoRepository(transactionContext *pgTransaction.TransactionContext) (*DemoRepository, error) {
	if transactionContext == nil {
		return nil, fmt.Errorf("transactionContext参数不能为nil")
	} else {
		return &DemoRepository{
			transactionContext: transactionContext,
		}, nil
	}
}
