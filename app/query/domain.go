package query

import (
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/model"
	"gorm.io/gorm"
)

func ListDomainUsers(db *gorm.DB, domain *model.Domain) *gorm.DB {
	return db.Table("domain_users").
		Select("domain_users.created_at, domain_users.updated_at, "+
			"domain_users.domain_id, domain_users.user_id, "+
			"domain_users.id, domain_users.role as domain_role, "+
			"users.username, users.gravatar").
		Joins("JOIN users ON domain_users.user_id = users.id").
		Where("domain_users.domain_id = ?", domain.Id)
}

func SearchDomainCandidates(
	db *gorm.DB, domainId uuid.UUID, query string,
) *gorm.DB {
	return db.Table("users").
		Select("users.id, users.username, users.gravatar, "+
			"domain_users.role as domain_role, domain_users.domain_id").
		Joins("LEFT OUTER JOIN domain_users ON "+
			"users.id = domain_users.user_id").
		Where("users.username ILIKE ?", query).
		Where(db.Where("domain_users.domain_id = ?", domainId).
			Or("domain_users.domain_id IS NULL"))
}
