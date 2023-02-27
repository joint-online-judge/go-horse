package query

import (
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/pkg/convert"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func GetDomain(
	domain string,
) (domainModel model.Domain, err error) {
	var query model.Domain
	if domainId, err := uuid.Parse(domain); err != nil {
		query.Url = domain
	} else {
		query.Id = domainId
	}
	domainModel, err = GetObjTo[model.Domain](&query)
	if err != nil {
		return
	}
	return
}

func GetDomainId(domain string) (uuid.UUID, error) {
	var query model.Domain
	if domainId, err := uuid.Parse(domain); err != nil {
		query.Url = domain
	} else {
		query.Id = domainId
	}
	err := db.Select("id").First(&query).Error
	return query.Id, err
}

func CreateDomain(
	domainCreate schema.DomainCreate,
	user *schema.User,
) (domain model.Domain, err error) {
	owner := model.User{Id: user.Id}
	err = convert.Update(&domain, domainCreate)
	if err != nil {
		return
	}
	domain.Owner = owner
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&domain).Error; err != nil {
			return err
		}
		logrus.Infof("create domain: %+v", domain)
		domainUser := model.DomainUser{
			Domain: domain,
			User:   owner,
			Role:   string(schema.ROOT),
		}
		if err := tx.Create(&domainUser).Error; err != nil {
			return err
		}
		logrus.Infof("create domain user: %+v", domainUser)
		// TODO: create domain roles with permissions
		return nil
	})
	return domain, err
}

func ListDomainUsers(domain *model.Domain, pagination schema.Pagination) (
	[]schema.UserWithDomainRole, int64, error,
) {
	statement := db.Table("domain_users").
		Select("domain_users.created_at, domain_users.updated_at, "+
			"domain_users.domain_id, domain_users.user_id, "+
			"domain_users.id, domain_users.role as domain_role, "+
			"users.username, users.gravatar").
		Joins("JOIN users ON domain_users.user_id = users.id").
		Where("domain_users.domain_id = ?", domain.Id)
	return ListObjs[schema.UserWithDomainRole](statement, pagination)
}

func AddDomainUser(domainId uuid.UUID, user model.User, role string) (
	u schema.UserWithDomainRole, err error,
) {
	model := model.DomainUser{
		Domain: model.Domain{Id: domainId},
		User:   user,
		Role:   role,
	}
	u, err = CreateObj[schema.UserWithDomainRole](&model)
	if err != nil {
		return
	}
	u.DomainRole = &model.Role
	u.Gravatar = &user.Gravatar
	u.Id = user.Id
	u.Username = user.Username
	return
}

func SearchDomainCandidates(
	domainId uuid.UUID, query string, pagination schema.Pagination,
) ([]schema.UserWithDomainRole, int64, error) {
	statement := db.Table("users").
		Select("users.id, users.username, users.gravatar, "+
			"domain_users.role as domain_role, domain_users.domain_id").
		Joins("LEFT OUTER JOIN domain_users ON "+
			"users.id = domain_users.user_id").
		Where("users.username ILIKE ?", query).
		Where(db.Where("domain_users.domain_id = ?", domainId).
			Or("domain_users.domain_id IS NULL"))
	return ListObjs[schema.UserWithDomainRole](statement, pagination)
}
