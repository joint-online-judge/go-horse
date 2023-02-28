package platform

import (
	_ "github.com/joint-online-judge/go-horse/pkg/config" // load config

	"github.com/joint-online-judge/go-horse/app/service"
	"github.com/joint-online-judge/go-horse/platform/auth"
	"github.com/joint-online-judge/go-horse/platform/cache"
	"github.com/joint-online-judge/go-horse/platform/db"
	"github.com/joint-online-judge/go-horse/platform/error"
	"github.com/joint-online-judge/go-horse/platform/queue"
	"github.com/joint-online-judge/go-horse/platform/storage"
)

func Bootstrap() {
	error.ConnectRollbar()
	db.ConnectPostgres()
	auth.Init()
	cache.ConnectRedis()
	storage.ConnectMinio()
	queue.ConnectAsynq()
	service.NewDB(db.DB)
}
