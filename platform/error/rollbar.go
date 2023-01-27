package error

import (
	"github.com/joint-online-judge/go-horse/pkg/configs"
	"github.com/rollbar/rollbar-go"
	"github.com/sirupsen/logrus"
)

func ConnectRollbar() {
	rollbar.SetToken(configs.Conf.RollbarAccessToken)
	rollbar.SetLogger(logrus.StandardLogger())
}
