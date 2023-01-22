package handlers

import (
	"context"

	"github.com/joint-online-judge/go-horse/handlers/utils"
	"github.com/joint-online-judge/go-horse/model"
	"github.com/joint-online-judge/go-horse/types"
	log "github.com/sirupsen/logrus"
)

func (s *ApiV1) V1CreateDomain(
	ctx context.Context,
	request types.V1CreateDomainRequestObject,
) (any, error) {
	b := request.Body
	log.Infof("request.Body: %v", b)
	return nil, nil
}

func (s *ApiV1) V1ListDomains(
	ctx context.Context,
	request types.V1ListDomainsRequestObject,
) (any, error) {
	return utils.NewListResp[model.Domain, types.Domain]()
}
