package platform

import (
	_ "github.com/joint-online-judge/go-horse/pkg/configs" // load config
	"github.com/joint-online-judge/go-horse/platform/cache"
	"github.com/joint-online-judge/go-horse/platform/db"
	"github.com/joint-online-judge/go-horse/platform/error"
	"github.com/joint-online-judge/go-horse/platform/queue"
	"github.com/joint-online-judge/go-horse/platform/storage"
)

func Bootstrap() {
	error.ConnectRollbar()
	db.ConnectPostgres()
	cache.ConnectRedis()
	queue.ConnectAsyncq()
	storage.ConnectMinio()
}
