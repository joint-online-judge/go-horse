package query

import (
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func GetDomain(
	domain string,
) (domainModel model.Domain, err error) {
	var query model.Domain
	if domainID, err := uuid.Parse(domain); err != nil {
		query.URL = domain
	} else {
		query.ID = domainID
	}
	domainModel, err = GetObj[model.Domain](&query)
	if err != nil {
		return
	}
	return
}

func GetDomainId(domain string) (uuid.UUID, error) {
	var query model.Domain
	if domainID, err := uuid.Parse(domain); err != nil {
		query.URL = domain
	} else {
		query.ID = domainID
	}
	err := db.Select("id").First(&query).Error
	return query.ID, err
}

func CreateDomain(
	domainCreate *schema.DomainCreate,
	user *schema.User,
) (any, error) {
	owner := model.User{ID: user.ID}
	domain := model.Domain{
		Owner:    owner,
		URL:      *domainCreate.Url,
		Name:     domainCreate.Name,
		Gravatar: *domainCreate.Gravatar,
		Bulletin: *domainCreate.Bulletin,
		Hidden:   *domainCreate.Hidden,
		Group:    *domainCreate.Group,
	}
	if err := db.Transaction(func(tx *gorm.DB) error {
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
	}); err != nil {
		return nil, err
	}
	return ConvertTo[schema.Domain](&domain)
}

func ListDomainUsers(domainId uuid.UUID, pagination schema.Pagination) (
	[]schema.UserWithDomainRole, int64, error,
) {
	statement := db.Table("domain_users").
		Select("domain_users.created_at, domain_users.updated_at, "+
			"domain_users.domain_id, domain_users.user_id, "+
			"domain_users.id, domain_users.role as domain_role, "+
			"users.username, users.gravatar").
		Joins("JOIN users ON domain_users.user_id = users.id").
		Where("domain_users.domain_id = ?", domainId)
	return ListObjs[schema.UserWithDomainRole](statement, pagination)
}

func AddDomainUser(domainId uuid.UUID, user schema.User, role string) (
	u schema.UserWithDomainRole, err error,
) {
	model := model.DomainUser{
		Domain: model.Domain{ID: domainId},
		User:   model.User{ID: user.ID},
		Role:   role,
	}
	u, err = CreateObj[schema.UserWithDomainRole](&model)
	if err != nil {
		return
	}
	u.DomainRole = &model.Role
	u.Gravatar = user.Gravatar
	u.ID = user.ID
	u.Username = user.Username
	return
}

func SearchDomainCandidates(
	domainId uuid.UUID, query string, pagination schema.Pagination,
) ([]schema.UserWithDomainRole, error) {
	statement := db.Table("users").
		Select("users.id, users.username, users.gravatar, "+
			"domain_users.role as domain_role, domain_users.domain_id").
		Joins("LEFT OUTER JOIN domain_users ON "+
			"users.id = domain_users.user_id").
		Where("users.username ILIKE ?", query).
		Where(db.Where("domain_users.domain_id = ?", domainId).
			Or("domain_users.domain_id IS NULL"))
	res, _, err := ListObjs[schema.UserWithDomainRole](statement, pagination)
	return res, err
}
