// Package schema provides primitives to interact with the openapi HTTP API.
package schema

import (
	"time"

	"github.com/google/uuid"
)

// Record defines model for Record.
type Record struct {
	CreatedAt *time.Time   `json:"createdAt,omitempty"`
	DomainId  uuid.UUID    `json:"domainId"`
	Id        uuid.UUID    `json:"id"`
	JudgedAt  *time.Time   `json:"judgedAt,omitempty"`
	Language  string       `json:"language"`
	MemoryKb  *int         `json:"memoryKb,omitempty"`
	Score     *int         `json:"score,omitempty"`
	State     *RecordState `json:"state,omitempty"`
	TimeMs    *int         `json:"timeMs,omitempty"`
	UpdatedAt *time.Time   `json:"updatedAt,omitempty"`
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
	Cases           *[]RecordCase `json:"cases,omitempty"`
	CommitId        *string       `json:"commitId,omitempty"`
	CommitterId     *uuid.UUID    `json:"committerId,omitempty"`
	CreatedAt       *time.Time    `json:"createdAt,omitempty"`
	DomainId        uuid.UUID     `json:"domainId"`
	Id              uuid.UUID     `json:"id"`
	JudgedAt        *time.Time    `json:"judgedAt,omitempty"`
	JudgerId        *uuid.UUID    `json:"judgerId,omitempty"`
	Language        string        `json:"language"`
	MemoryKb        *int          `json:"memoryKb,omitempty"`
	ProblemConfigId *uuid.UUID    `json:"problemConfigId,omitempty"`
	ProblemId       *uuid.UUID    `json:"problemId,omitempty"`
	ProblemSetId    *uuid.UUID    `json:"problemSetId,omitempty"`
	Score           *int          `json:"score,omitempty"`
	State           *RecordState  `json:"state,omitempty"`
	TaskId          *uuid.UUID    `json:"taskId,omitempty"`
	TimeMs          *int          `json:"timeMs,omitempty"`
	UpdatedAt       *time.Time    `json:"updatedAt,omitempty"`
}

// RecordListDetail defines model for RecordListDetail.
type RecordListDetail struct {
	CommitterId       *uuid.UUID   `json:"committerId,omitempty"`
	CommitterUsername *string      `json:"committerUsername,omitempty"`
	CreatedAt         *time.Time   `json:"createdAt,omitempty"`
	DomainId          uuid.UUID    `json:"domainId"`
	Id                uuid.UUID    `json:"id"`
	JudgedAt          *time.Time   `json:"judgedAt,omitempty"`
	Language          string       `json:"language"`
	MemoryKb          *int         `json:"memoryKb,omitempty"`
	ProblemId         *uuid.UUID   `json:"problemId,omitempty"`
	ProblemSetId      *uuid.UUID   `json:"problemSetId,omitempty"`
	ProblemSetTitle   *string      `json:"problemSetTitle,omitempty"`
	ProblemTitle      *string      `json:"problemTitle,omitempty"`
	Score             *int         `json:"score,omitempty"`
	State             *RecordState `json:"state,omitempty"`
	TimeMs            *int         `json:"timeMs,omitempty"`
	UpdatedAt         *time.Time   `json:"updatedAt,omitempty"`
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
	CreatedAt time.Time `json:"createdAt"`
	Id        uuid.UUID `json:"id"`

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
	ProblemSet *uuid.UUID `form:"problemSet,omitempty" json:"problemSet,omitempty"`

	// Problem problem id
	Problem *uuid.UUID `form:"problem,omitempty" json:"problem,omitempty"`

	// SubmitterId submitter uid
	SubmitterId *uuid.UUID `form:"submitterId,omitempty" json:"submitterId,omitempty"`
	Pagination
}
