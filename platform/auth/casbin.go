package auth

import (
	_ "embed"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/joint-online-judge/go-horse/platform/db"
	"github.com/sirupsen/logrus"
)

var (
	Adapter  *gormadapter.Adapter
	Enforcer *casbin.Enforcer
	//go:embed rbac_with_domains_model.conf
	rbacWithDomainsModelConfText string
)

func Init() {
	var err error
	Adapter, err = gormadapter.NewAdapterByDB(db.DB)
	if err != nil {
		logrus.Fatalf("failed to create casbin adapter: %+v", err)
	}
	m, err := model.NewModelFromString(rbacWithDomainsModelConfText)
	if err != nil {
		logrus.Fatalf("failed to create casbin model: %+v", err)
	}
	Enforcer, err = casbin.NewEnforcer(m, Adapter)
	if err != nil {
		logrus.Fatalf("failed to create casbin enforcer: %+v", err)
	}
}
