package transform

import (
	"github.com/linmadan/mmm-newdemo/pkg/domain"
	"github.com/linmadan/mmm-newdemo/pkg/infrastructure/pg/models"
)

func TransformToDemoDomainModelFromPgModels(demoModel *models.Demo) (*domain.Demo, error) {
	return &domain.Demo{
		DemoId:   demoModel.DemoId,
		DemoName: demoModel.DemoName,
	}, nil
}
