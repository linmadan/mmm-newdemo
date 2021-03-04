package command

import (
	"fmt"

	"github.com/beego/beego/v2/core/validation"
)

type CreateDemoCommand struct {
	// 例子名称
	DemoName string `json:"demoName" valid:"Required"`
}

func (createDemoCommand *CreateDemoCommand) ValidateCommand() error {
	valid := validation.Validation{}
	b, err := valid.Valid(createDemoCommand)
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
