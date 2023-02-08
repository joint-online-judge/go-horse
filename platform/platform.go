package platform

import (
	_ "github.com/joint-online-judge/go-horse/pkg/configs" // load config
	"github.com/joint-online-judge/go-horse/platform/db"
	"github.com/joint-online-judge/go-horse/platform/error"
)

func Bootstrap() {
	error.ConnectRollbar()
	db.ConnectPostgres()
	// cache.ConnectRedis()
	// queue.ConnectAsyncq()
	// storage.ConnectMinio()
}
