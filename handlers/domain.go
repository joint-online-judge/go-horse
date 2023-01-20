package handlers

import (
	"context"

	"github.com/joint-online-judge/go-horse/types"
	log "github.com/sirupsen/logrus"
)

func (s *ApiV1) V1CreateDomain(c context.Context, request types.V1CreateDomainRequestObject) (types.V1CreateDomainResponseObject, error) {
	log.Infof("%v", request.Body)
	if err, ok := ValidateStructs(request.Body); !ok {
		return types.V1CreateDomain422JSONResponse{Detail: &err}, nil
	}
	return types.NewV1CreateDomain200JSONResponse(&types.Domain{}), nil
}
