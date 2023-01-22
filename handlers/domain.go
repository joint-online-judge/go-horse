package handlers

import (
	"context"

	"github.com/joint-online-judge/go-horse/types"
	log "github.com/sirupsen/logrus"
)

func (s *ApiV1) V1CreateDomain(c context.Context, request types.V1CreateDomainRequestObject) (any, error) {
	b := request.Body
	log.Infof("request.Body: %v", b)
	// FIXME: get enough parameters
	// domain := model.Domain{Name: b.Name, URL: *b.Url, Gravatar: *b.Gravatar}
	// err := dao.Domain.WithContext(c).Create(&domain)
	// if err != nil {
	// 	return nil, err
	// }
	// log.Infof("domain: %v", domain)
	return nil, nil
}
