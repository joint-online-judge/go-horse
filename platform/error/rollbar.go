package error

import (
	"github.com/joint-online-judge/go-horse/pkg/config"
	"github.com/rollbar/rollbar-go"
	"github.com/sirupsen/logrus"
)

func ConnectRollbar() {
	rollbar.SetToken(config.Conf.RollbarAccessToken)
	rollbar.SetLogger(logrus.StandardLogger())
}
