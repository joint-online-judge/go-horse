package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/query"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/pkg/convert"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type domainImpl struct {
	c *fiber.Ctx
}

func Domain(c *fiber.Ctx) *domainImpl {
	return &domainImpl{
		c: c,
	}
}

func (s *domainImpl) GetDomain(domain string) (
	domainModel model.Domain, err error,
) {
	if domainId, err := uuid.Parse(domain); err != nil {
		domainModel.Url = domain
	} else {
		domainModel.Id = domainId
	}
	err = db.First(&domainModel).Error
	return
}

// GetCurrentDomain Assume we require current domain to be non-null
// when calling this method
func (s *domainImpl) GetCurrentDomain() (*model.Domain, error) {
	domain := s.c.Locals("domain")
	if domain == nil {
		return nil, schema.NewBizError(
			schema.DomainNotFoundError, "cannot find current domain",
		)
	}
	return domain.(*model.Domain), nil
}

func (s *domainImpl) ListDomains(params schema.ListDomainsParams) (
	schema.ListResp[schema.Domain], error,
) {
	// TODO: filter by domain users
	objs, count, err := ListObjs[schema.Domain](
		db.Model(model.Domain{}), params.Pagination,
	)
	return schema.NewListResp(count, objs), err
}

func (s *domainImpl) CreateDomain(domainCreate schema.DomainCreate) (
	domain model.Domain, err error,
) {
	user := Auth(s.c).JWTUser()
	// TODO: verify input values
	owner := model.User{Id: user.Id}
	err = convert.Update(&domain, domainCreate)
	if err != nil {
		return
	}
	domain.Owner = &owner
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&domain).Error; err != nil {
			return err
		}
		logrus.Debugf("create domain: %+v", domain)
		domainUser := model.DomainUser{
			Domain: domain,
			User:   owner,
			Role:   string(schema.ROOT),
		}
		if err := tx.Create(&domainUser).Error; err != nil {
			return err
		}
		logrus.Debugf("create domain user: %+v", domainUser)
		// TODO: create domain roles with permissions
		return nil
	})
	return
}

func (s *domainImpl) UpdateDomain(domainEdit schema.DomainEdit) (
	domain *model.Domain, err error,
) {
	domain, err = s.GetCurrentDomain()
	if err != nil {
		return
	}
	if err = convert.Update(domain, domainEdit); err != nil {
		return
	}
	logrus.Debugf("update domain to: %+v", domain)
	err = db.Save(domain).Error
	return
}

func (s *domainImpl) SearchDomainCandidates(
	ordering *string, searchQuery string,
) (resp schema.ListResp[schema.UserWithDomainRole], err error) {
	domain, err := s.GetCurrentDomain()
	if err != nil {
		return
	}
	pagination := schema.Pagination{
		Ordering: ordering,
		Offset:   nil,
		Limit:    nil,
	}
	objs, count, err := ListObjs[schema.UserWithDomainRole](
		query.SearchDomainCandidates(
			db, domain.Id, searchQuery,
		), pagination)
	return schema.NewListResp(count, objs), err
}

func (s *domainImpl) ListDomainUsers(
	pagination schema.Pagination,
) (resp schema.ListResp[schema.UserWithDomainRole], err error) {
	domain, err := s.GetCurrentDomain()
	if err != nil {
		return
	}
	objs, count, err := ListObjs[schema.UserWithDomainRole](
		query.ListDomainUsers(db, domain), pagination,
	)
	return schema.NewListResp(count, objs), err
}

func (s *domainImpl) AddDomainUser(
	user string, role *string,
) (resp schema.UserWithDomainRole, err error) {
	domain, err := s.GetCurrentDomain()
	if err != nil {
		return
	}
	userModel, err := User(s.c).GetUser(user)
	if err != nil {
		return
	}
	if role == nil {
		role = schema.Pointer(string(schema.USER))
	}
	model := model.DomainUser{
		Domain: model.Domain{Id: domain.Id},
		User:   userModel,
		Role:   *role,
	}
	err = db.Create(&model).Error
	if err != nil {
		return
	}
	resp.DomainRole = role
	resp.Gravatar = &userModel.Gravatar
	resp.Id = userModel.Id
	resp.Username = userModel.Username
	return
}
