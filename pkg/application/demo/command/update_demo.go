package command

import (
	"fmt"

	"github.com/beego/beego/v2/core/validation"
)

type UpdateDemoCommand struct {
	// 例子id
	DemoId int64 `json:"demoId" valid:"Required"`
	// 例子名称
	DemoName string `json:"demoName,omitempty"`
}

func (updateDemoCommand *UpdateDemoCommand) ValidateCommand() error {
	valid := validation.Validation{}
	b, err := valid.Valid(updateDemoCommand)
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
