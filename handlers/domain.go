package handlers

import (
	"context"

	"github.com/joint-online-judge/go-horse/database"
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
	var domains []types.Domain
	db := database.DB
	err := db.Model(&model.Domain{}).Find(&domains).Error
	if err != nil {
		return nil, err
	}
	return types.ListResp[types.Domain]{Count: 0, Results: domains}, nil
}
