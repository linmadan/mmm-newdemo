package log

import (
	"github.com/linmadan/egglib-go/log"
	"github.com/linmadan/egglib-go/log/logrus"
	"github.com/linmadan/mmm-newdemo/pkg/constant"
)

var Logger log.Logger

func init() {
	Logger = logrus.NewLogrusLogger()
	Logger.SetServiceName(constant.SERVICE_NAME)
	Logger.SetLevel(constant.LOG_LEVEL)
}
