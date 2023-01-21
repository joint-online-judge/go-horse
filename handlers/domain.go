package handlers

import (
	"context"

	"github.com/joint-online-judge/go-horse/types"
	log "github.com/sirupsen/logrus"
)

func (s *ApiV1) V1CreateDomain(c context.Context, request types.V1CreateDomainRequestObject) (interface{}, error) {
	log.Infof("%v", request.Body)
	domain := types.Domain{}
	return domain, nil
}
