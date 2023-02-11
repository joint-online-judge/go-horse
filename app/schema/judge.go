// Package schema provides primitives to interact with the openapi HTTP API.
package schema

import (
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// JudgerClaim defines model for JudgerClaim.
type JudgerClaim struct {
	TaskId string `json:"taskId"`
}

// JudgerCreate defines model for JudgerCreate.
type JudgerCreate struct {
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
	Username *string `json:"username,omitempty"`
}

// JudgerCredentials defines model for JudgerCredentials.
type JudgerCredentials struct {
	ProblemConfigCommitId string `json:"problemConfigCommitId"`
	ProblemConfigRepoName string `json:"problemConfigRepoName"`
	RecordCommitId        string `json:"recordCommitId"`
	RecordRepoName        string `json:"recordRepoName"`
}

// JudgerDetail defines model for JudgerDetail.
type JudgerDetail struct {
	CreatedAt *time.Time         `json:"createdAt,omitempty"`
	Gravatar  *string            `json:"gravatar,omitempty"`
	Id        openapi_types.UUID `json:"id"`
	IsActive  *bool              `json:"isActive,omitempty"`
	IsAlive   bool               `json:"isAlive"`
	Role      *string            `json:"role,omitempty"`
	UpdatedAt *time.Time         `json:"updatedAt,omitempty"`
	Username  string             `json:"username"`
}

// JudgerDetailList defines model for JudgerDetailList.
type JudgerDetailList struct {
	Count   *int            `json:"count,omitempty"`
	Results *[]JudgerDetail `json:"results,omitempty"`
}
