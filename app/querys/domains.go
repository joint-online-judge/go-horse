package querys

import (
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/models"
	"github.com/joint-online-judge/go-horse/app/schemas"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func GetDomain[Schema any](
	domain string,
) (domainModel models.Domain, domainSchema Schema, err error) {
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
	domainSchema, err = ConvertTo[Schema](domainModel)
	return
}

func GetDomainId(domain string) (string, error) {
	var query models.Domain
	if domainID, err := uuid.Parse(domain); err != nil {
		query.URL = domain
	} else {
		query.ID = domainID
	}
	err := DB.Select("id").First(&query).Error
	return query.ID.String(), err
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

func ListDomainUsers(domainId string, pagination schemas.Pagination) ([]schemas.UserWithDomainRole, int64, error) {
	var results []schemas.UserWithDomainRole
	u := DB.Table("domain_users").
		Select("domain_users.created_at, domain_users.updated_at, domain_users.domain_id, domain_users.user_id, domain_users.id, domain_users.role as domain_role, users.username, users.gravatar").
		Joins("JOIN users ON domain_users.user_id = users.id").
		Where("domain_users.domain_id = ?", domainId)
	var count int64
	if err := u.Count(&count).Error; err != nil {
		return nil, 0, errors.Wrapf(err, "failed get %T count", models.DomainUser{})
	}
	err := u.Scopes(Paginate(pagination)).Scan(&results).Error
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}
	return results, count, nil
}
