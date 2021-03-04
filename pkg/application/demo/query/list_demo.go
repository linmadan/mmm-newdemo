package query

import (
	"fmt"

	"github.com/beego/beego/v2/core/validation"
)

type ListDemoQuery struct {
	// 查询偏离量
	Offset int `json:"offset,omitempty"`
	// 查询限制
	Limit int `json:"limit,omitempty"`
}

func (listDemoQuery *ListDemoQuery) ValidateQuery() error {
	valid := validation.Validation{}
	b, err := valid.Valid(listDemoQuery)
	if err != nil {
		return err
	}
	if !b {
		for _, validErr := range valid.Errors {
			return fmt.Errorf("%s  %s", validErr.Key, validErr.Message)
		}
	}
	return nil
}
