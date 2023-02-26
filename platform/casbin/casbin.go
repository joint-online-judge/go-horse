package casbin

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/joint-online-judge/go-horse/pkg/config"
	"github.com/joint-online-judge/go-horse/platform/db"
	"github.com/sirupsen/logrus"
)

var Adapter *gormadapter.Adapter
var Enforcer *casbin.Enforcer

func Init() {
	var err error
	modelFilePath := config.Conf.CasbinModelFilePath
	logrus.Infof("Model file: %v", modelFilePath)
	Adapter, err = gormadapter.NewAdapterByDB(db.DB)
	if err != nil {
		logrus.Fatalf("failed to create casbin adapter: %+v", err)
	}
	Enforcer, err = casbin.NewEnforcer(modelFilePath, Adapter)
	if err != nil {
		logrus.Fatalf("failed to create casbin enforcer: %+v", err)
	}
}
