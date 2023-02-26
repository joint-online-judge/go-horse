package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/query"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/pkg/convert"
	"github.com/sirupsen/logrus"
)

type domainImpl struct {
	c *fiber.Ctx
}

func Domain(c *fiber.Ctx) *domainImpl {
	return &domainImpl{
		c: c,
	}
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
	objs, count, err := query.ListObjsByType[
		model.Domain, schema.Domain,
	](params.Pagination)
	return schema.NewListResp(count, objs), err
}

func (s *domainImpl) CreateDomain(domainCreate schema.DomainCreate) (
	model.Domain, error,
) {
	// TODO: verify input values
	user := Auth(s.c).JWTUser()
	return query.CreateDomain(domainCreate, user)
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
	logrus.Infof("update domain to: %+v", domain)
	err = query.SaveObj(domain)
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
	objs, count, err := query.SearchDomainCandidates(
		domain.ID, searchQuery, pagination,
	)
	return schema.NewListResp(count, objs), err
}

func (s *domainImpl) ListDomainUsers(
	pagination schema.Pagination,
) (resp schema.ListResp[schema.UserWithDomainRole], err error) {
	domain, err := s.GetCurrentDomain()
	if err != nil {
		return
	}
	objs, count, err := query.ListDomainUsers(domain, pagination)
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
	return query.AddDomainUser(domain.ID, userModel, *role)
}
