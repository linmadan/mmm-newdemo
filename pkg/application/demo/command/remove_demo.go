package command

import (
	"fmt"

	"github.com/beego/beego/v2/core/validation"
)

type RemoveDemoCommand struct {
	// 例子id
	DemoId int64 `json:"demoId" valid:"Required"`
}

func (removeDemoCommand *RemoveDemoCommand) ValidateCommand() error {
	valid := validation.Validation{}
	b, err := valid.Valid(removeDemoCommand)
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
