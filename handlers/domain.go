package handlers

import (
	"context"

	// "github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/types"
	log "github.com/sirupsen/logrus"
)

func (s *ApiV1) V1CreateDomain(c context.Context, request types.V1CreateDomainRequestObject) (types.V1CreateDomainResponseObject, error) {
	log.Infof("%v", request.Body)
	return types.V1CreateDomain200JSONResponse{BizError: types.BizError{ErrorCode: "Success"}, Data: &types.Domain{}}, nil
}
