// Package schema provides primitives to interact with the openapi HTTP API.
package schema

import (
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// Record defines model for Record.
type Record struct {
	CreatedAt *time.Time         `json:"createdAt,omitempty"`
	DomainId  openapi_types.UUID `json:"domainId"`
	Id        openapi_types.UUID `json:"id"`
	JudgedAt  *time.Time         `json:"judgedAt,omitempty"`
	Language  string             `json:"language"`
	MemoryKb  *int               `json:"memoryKb,omitempty"`
	Score     *int               `json:"score,omitempty"`
	State     *RecordState       `json:"state,omitempty"`
	TimeMs    *int               `json:"timeMs,omitempty"`
	UpdatedAt *time.Time         `json:"updatedAt,omitempty"`
}

// RecordCase defines model for RecordCase.
type RecordCase struct {
	MemoryKb   *int              `json:"memoryKb,omitempty"`
	ReturnCode *int              `json:"returnCode,omitempty"`
	Score      *int              `json:"score,omitempty"`
	State      *RecordCaseResult `json:"state,omitempty"`
	Stderr     *string           `json:"stderr,omitempty"`
	Stdout     *string           `json:"stdout,omitempty"`
	TimeMs     *int              `json:"timeMs,omitempty"`
}

// RecordCaseSubmit defines model for RecordCaseSubmit.
type RecordCaseSubmit struct {
	MemoryKb   *int `json:"memoryKb,omitempty"`
	ReturnCode *int `json:"returnCode,omitempty"`
	Score      *int `json:"score,omitempty"`

	// State An enumeration.
	State  *RecordCaseResult `json:"state,omitempty"`
	Stderr *string           `json:"stderr,omitempty"`
	Stdout *string           `json:"stdout,omitempty"`
	TimeMs *int              `json:"timeMs,omitempty"`
}

// RecordDetail defines model for RecordDetail.
type RecordDetail struct {
	Cases           *[]RecordCase       `json:"cases,omitempty"`
	CommitId        *string             `json:"commitId,omitempty"`
	CommitterId     *openapi_types.UUID `json:"committerId,omitempty"`
	CreatedAt       *time.Time          `json:"createdAt,omitempty"`
	DomainId        openapi_types.UUID  `json:"domainId"`
	Id              openapi_types.UUID  `json:"id"`
	JudgedAt        *time.Time          `json:"judgedAt,omitempty"`
	JudgerId        *openapi_types.UUID `json:"judgerId,omitempty"`
	Language        string              `json:"language"`
	MemoryKb        *int                `json:"memoryKb,omitempty"`
	ProblemConfigId *openapi_types.UUID `json:"problemConfigId,omitempty"`
	ProblemId       *openapi_types.UUID `json:"problemId,omitempty"`
	ProblemSetId    *openapi_types.UUID `json:"problemSetId,omitempty"`
	Score           *int                `json:"score,omitempty"`
	State           *RecordState        `json:"state,omitempty"`
	TaskId          *openapi_types.UUID `json:"taskId,omitempty"`
	TimeMs          *int                `json:"timeMs,omitempty"`
	UpdatedAt       *time.Time          `json:"updatedAt,omitempty"`
}

// RecordListDetail defines model for RecordListDetail.
type RecordListDetail struct {
	CommitterId       *openapi_types.UUID `json:"committerId,omitempty"`
	CommitterUsername *string             `json:"committerUsername,omitempty"`
	CreatedAt         *time.Time          `json:"createdAt,omitempty"`
	DomainId          openapi_types.UUID  `json:"domainId"`
	Id                openapi_types.UUID  `json:"id"`
	JudgedAt          *time.Time          `json:"judgedAt,omitempty"`
	Language          string              `json:"language"`
	MemoryKb          *int                `json:"memoryKb,omitempty"`
	ProblemId         *openapi_types.UUID `json:"problemId,omitempty"`
	ProblemSetId      *openapi_types.UUID `json:"problemSetId,omitempty"`
	ProblemSetTitle   *string             `json:"problemSetTitle,omitempty"`
	ProblemTitle      *string             `json:"problemTitle,omitempty"`
	Score             *int                `json:"score,omitempty"`
	State             *RecordState        `json:"state,omitempty"`
	TimeMs            *int                `json:"timeMs,omitempty"`
	UpdatedAt         *time.Time          `json:"updatedAt,omitempty"`
}

// RecordListDetailList defines model for RecordListDetailList.
type RecordListDetailList struct {
	Count   *int                `json:"count,omitempty"`
	Results *[]RecordListDetail `json:"results,omitempty"`
}

// RecordPermission defines model for RecordPermission.
type RecordPermission struct {
	Code    *bool `json:"code,omitempty"`
	Detail  *bool `json:"detail,omitempty"`
	Judge   *bool `json:"judge,omitempty"`
	Rejudge *bool `json:"rejudge,omitempty"`
	View    *bool `json:"view,omitempty"`
}

// RecordPreview defines model for RecordPreview.
type RecordPreview struct {
	CreatedAt time.Time          `json:"createdAt"`
	Id        openapi_types.UUID `json:"id"`

	// State An enumeration.
	State RecordState `json:"state"`
}

// RecordSubmit defines model for RecordSubmit.
type RecordSubmit struct {
	JudgedAt *time.Time `json:"judgedAt,omitempty"`
	MemoryKb *int       `json:"memoryKb,omitempty"`
	Score    *int       `json:"score,omitempty"`

	// State An enumeration.
	State  *RecordState `json:"state,omitempty"`
	TimeMs *int         `json:"timeMs,omitempty"`
}

// ListRecordsInDomainParams defines parameters for ListRecordsInDomain.
type ListRecordsInDomainParams struct {
	// ProblemSet problem set id
	ProblemSet *openapi_types.UUID `form:"problemSet,omitempty" json:"problemSet,omitempty"`

	// Problem problem id
	Problem *openapi_types.UUID `form:"problem,omitempty" json:"problem,omitempty"`

	// SubmitterId submitter uid
	SubmitterId *openapi_types.UUID `form:"submitterId,omitempty" json:"submitterId,omitempty"`
	Pagination
}
