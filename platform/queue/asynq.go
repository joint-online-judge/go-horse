package queue

import (
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/joint-online-judge/go-horse/pkg/config"
	"github.com/sirupsen/logrus"
)

var Asynq *asynq.Client

func ConnectAsynq() {
	defer func() {
		if err := recover(); err != nil {
			logrus.Fatal(err)
		}
	}()
	conf := config.Conf
	config := asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%d", conf.RedisHost, conf.RedisPort),
		Password: conf.RedisPassword,
		DB:       conf.RedisDbIndex,
	}
	Asynq = asynq.NewClient(config)
	logrus.Debugf("asynq client: %+v", Asynq)
}
