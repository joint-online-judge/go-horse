package error

import (
	"github.com/joint-online-judge/go-horse/pkg/configs"
	"github.com/rollbar/rollbar-go"
)

func ConnectRollbar() {
	rollbar.SetToken(configs.Conf.RollbarAccessToken)
}
