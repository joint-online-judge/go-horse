package querys

import (
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/models"
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
