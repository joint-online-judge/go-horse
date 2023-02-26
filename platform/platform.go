package platform

import (
	_ "github.com/joint-online-judge/go-horse/pkg/config" // load config
	"github.com/joint-online-judge/go-horse/platform/casbin"
	"github.com/joint-online-judge/go-horse/platform/db"
	"github.com/joint-online-judge/go-horse/platform/error"
)

func Bootstrap() {
	error.ConnectRollbar()
	db.ConnectPostgres()
	casbin.Init()
	// cache.ConnectRedis()
	// queue.ConnectAsyncq()
	// storage.ConnectMinio()
}
