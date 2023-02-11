package queue

import (
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/joint-online-judge/go-horse/pkg/config"
	"github.com/sirupsen/logrus"
)

func ConnectAsyncq() {
	conf := config.Conf
	config := asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%d", conf.RedisHost, conf.RedisPort),
		Password: conf.RedisPassword,
		DB:       conf.RedisDbIndex,
	}
	client := asynq.NewClient(config)
	logrus.Infof("asyncq client: %+v", client)
	defer client.Close()
}
