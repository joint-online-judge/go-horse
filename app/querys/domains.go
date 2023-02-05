package querys

import (
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/models"
	"github.com/joint-online-judge/go-horse/app/schemas"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func GetDomain(
	domain string,
) (domainModel models.Domain, err error) {
	var query models.Domain
	if domainID, err := uuid.Parse(domain); err != nil {
		query.URL = domain
	} else {
		query.ID = domainID
	}
	domainModel, err = GetObj[models.Domain](&query)
	if err != nil {
		return
	}
	return
}

func GetDomainId(domain string) (uuid.UUID, error) {
	var query models.Domain
	if domainID, err := uuid.Parse(domain); err != nil {
		query.URL = domain
	} else {
		query.ID = domainID
	}
	err := DB.Select("id").First(&query).Error
	return query.ID, err
}

func CreateDomain(
	domainCreate *schemas.DomainCreate,
	user *schemas.User,
) (any, error) {
	owner := models.User{Base: models.Base{ID: user.Id}}
	domain := models.Domain{
		Owner:    owner,
		URL:      *domainCreate.Url,
		Name:     domainCreate.Name,
		Gravatar: *domainCreate.Gravatar,
		Bulletin: *domainCreate.Bulletin,
		Hidden:   *domainCreate.Hidden,
		Group_:   *domainCreate.Group,
	}
	if err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&domain).Error; err != nil {
			return err
		}
		logrus.Infof("create domain: %+v", domain)
		domainUser := models.DomainUser{
			Domain: domain,
			User:   owner,
			Role:   string(schemas.ROOT),
		}
		if err := tx.Create(&domainUser).Error; err != nil {
			return err
		}
		logrus.Infof("create domain user: %+v", domainUser)
		// TODO: create domain roles with premissions
		return nil
	}); err != nil {
		return nil, err
	}
	return ConvertTo[schemas.Domain](&domain)
}

func ListDomainUsers(domainId uuid.UUID, pagination schemas.Pagination) (
	[]schemas.UserWithDomainRole, int64, error,
) {
	statment := DB.Table("domain_users").
		Select("domain_users.created_at, domain_users.updated_at, "+
			"domain_users.domain_id, domain_users.user_id, "+
			"domain_users.id, domain_users.role as domain_role, "+
			"users.username, users.gravatar").
		Joins("JOIN users ON domain_users.user_id = users.id").
		Where("domain_users.domain_id = ?", domainId)
	return ListObjs[schemas.UserWithDomainRole](statment, pagination)
}

func AddDomainUser(domainId uuid.UUID, user schemas.User, role string) (
	schema schemas.UserWithDomainRole, err error,
) {
	model := models.DomainUser{
		Domain: models.Domain{Base: models.Base{ID: domainId}},
		User:   models.User{Base: models.Base{ID: user.Id}},
		Role:   role,
	}
	schema, err = CreateObj[schemas.UserWithDomainRole](&model)
	if err != nil {
		return
	}
	schema.DomainRole = &model.Role
	schema.Gravatar = user.Gravatar
	schema.Id = user.Id
	schema.Username = user.Username
	return
}

func SearchDomainCandidates(
	domainId uuid.UUID, query string, pagination schemas.Pagination,
) ([]schemas.UserWithDomainRole, error) {
	statment := DB.Table("users").
		Select("users.id, users.username, users.gravatar, "+
			"domain_users.role as domain_role, domain_users.domain_id").
		Joins("LEFT OUTER JOIN domain_users ON "+
			"users.id = domain_users.user_id").
		Where("users.username ILIKE ?", query).
		Where(DB.Where("domain_users.domain_id = ?", domainId).
			Or("domain_users.domain_id IS NULL"))
	res, _, err := ListObjs[schemas.UserWithDomainRole](statment, pagination)
	return res, err
}
