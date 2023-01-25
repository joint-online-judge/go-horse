package cache

import (
	"context"
	"fmt"

	"github.com/joint-online-judge/go-horse/pkg/configs"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

var RDB *redis.Client

func ConnectRedis() {
	conf := configs.Conf
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.RedisHost, conf.RedisPort),
		Password: conf.RedisPassword,
		DB:       conf.RedisDbIndex,
	})
	status := RDB.Ping(context.Background())
	if err := status.Err(); err != nil {
		log.Fatalf("failed to ping to redis: %+v", err)
	}
}
