package query

import (
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/schema"
	"gorm.io/gorm"
)

func GetDomain(
	db *gorm.DB, domain string,
) (domainModel model.Domain, err error) {
	if domainId, err := uuid.Parse(domain); err != nil {
		domainModel.Url = domain
	} else {
		domainModel.Id = domainId
	}
	err = db.First(&domainModel).Error
	return
}

func ListDomainUsers(
	db *gorm.DB, domain *model.Domain, pagination schema.Pagination,
) ([]schema.UserWithDomainRole, int64, error) {
	statement := db.Table("domain_users").
		Select("domain_users.created_at, domain_users.updated_at, "+
			"domain_users.domain_id, domain_users.user_id, "+
			"domain_users.id, domain_users.role as domain_role, "+
			"users.username, users.gravatar").
		Joins("JOIN users ON domain_users.user_id = users.id").
		Where("domain_users.domain_id = ?", domain.Id)
	return ListObjs[schema.UserWithDomainRole](statement, pagination)
}

func SearchDomainCandidates(
	db *gorm.DB, domainId uuid.UUID, query string, pagination schema.Pagination,
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
