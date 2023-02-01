package querys

import (
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/models"
	"github.com/joint-online-judge/go-horse/app/schemas"
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
		}
		if err := tx.Create(&domain).Error; err != nil {
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
