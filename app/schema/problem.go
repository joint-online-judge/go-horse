// Package schema provides primitives to interact with the openapi HTTP API.
package schema

import (
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/google/uuid"
)

// Problem defines model for Problem.
type Problem struct {
	DomainId openapi_types.UUID `json:"domainId"`

	// Hidden is the problem hidden
	Hidden         *bool               `json:"hidden,omitempty"`
	Id             uuid.UUID           `json:"id"`
	NumAccept      *int                `json:"numAccept,omitempty"`
	NumSubmit      *int                `json:"numSubmit,omitempty"`
	OwnerId        *openapi_types.UUID `json:"ownerId,omitempty"`
	ProblemGroupId *openapi_types.UUID `json:"problemGroupId,omitempty"`

	// Title title of the problem
	Title string `json:"title"`

	// Url (unique) url of the domain
	Url *string `json:"url,omitempty"`
}

// ProblemClone defines model for ProblemClone.
type ProblemClone struct {
	// FromDomain url or id of the domain to clone from
	FromDomain string `json:"fromDomain"`

	// NewGroup whether to create new problem group
	NewGroup *bool    `json:"newGroup,omitempty"`
	Problems []string `json:"problems"`
}

// ProblemConfigDataDetail defines model for ProblemConfigDataDetail.
type ProblemConfigDataDetail struct {
	CommitId      *string             `json:"commitId,omitempty"`
	CommitMessage *string             `json:"commitMessage,omitempty"`
	CommitterId   *openapi_types.UUID `json:"committerId,omitempty"`
	CreatedAt     *time.Time          `json:"createdAt,omitempty"`
	Data          ProblemConfigJson   `json:"data"`
	DataVersion   *int                `json:"dataVersion,omitempty"`
	Id            uuid.UUID           `json:"id"`
	UpdatedAt     *time.Time          `json:"updatedAt,omitempty"`
}

// ProblemConfigDetail defines model for ProblemConfigDetail.
type ProblemConfigDetail struct {
	CommitId      *string             `json:"commitId,omitempty"`
	CommitMessage *string             `json:"commitMessage,omitempty"`
	CommitterId   *openapi_types.UUID `json:"committerId,omitempty"`
	CreatedAt     *time.Time          `json:"createdAt,omitempty"`
	DataVersion   *int                `json:"dataVersion,omitempty"`
	Id            uuid.UUID           `json:"id"`
	UpdatedAt     *time.Time          `json:"updatedAt,omitempty"`
}

// ProblemConfigDetailList defines model for ProblemConfigDetailList.
type ProblemConfigDetailList struct {
	Count   *int                   `json:"count,omitempty"`
	Results *[]ProblemConfigDetail `json:"results,omitempty"`
}

// Case defines model for Case.
type Case struct {
	Category          *string   `json:"category,omitempty"`
	ExecuteArgs       *[]string `json:"executeArgs,omitempty"`
	ExecuteFiles      *[]string `json:"executeFiles,omitempty"`
	ExecuteInputFile  *string   `json:"executeInputFile,omitempty"`
	ExecuteOutputFile *string   `json:"executeOutputFile,omitempty"`
	IgnoreWhitespace  *bool     `json:"ignoreWhitespace,omitempty"`
	Memory            *string   `json:"memory,omitempty"`
	Score             *int      `json:"score,omitempty"`
	Time              *string   `json:"time,omitempty"`
}

// LanguageDefault defines model for LanguageDefault.
type LanguageDefault struct {
	CaseDefault  *Case     `json:"caseDefault,omitempty"`
	Cases        *[]Case   `json:"cases,omitempty"`
	CompileArgs  *[]string `json:"compileArgs,omitempty"`
	CompileFiles *[]string `json:"compileFiles,omitempty"`
}

// Language defines model for Language.
type Language struct {
	CaseDefault  *Case     `json:"caseDefault,omitempty"`
	Cases        *[]Case   `json:"cases,omitempty"`
	CompileArgs  *[]string `json:"compileArgs,omitempty"`
	CompileFiles *[]string `json:"compileFiles,omitempty"`
	Name         string    `json:"name"`
}

// ProblemConfigJson defines model for ProblemConfigJson.
type ProblemConfigJson struct {
	LanguageDefault *LanguageDefault `json:"languageDefault,omitempty"`
	Languages       []Language       `json:"languages"`
}

// ProblemCreate defines model for ProblemCreate.
type ProblemCreate struct {
	// Content content of the problem
	Content *string `json:"content,omitempty"`

	// Hidden is the problem hidden
	Hidden *bool `json:"hidden,omitempty"`

	// Title title of the problem
	Title string `json:"title"`

	// Url (unique) url of the domain
	Url *string `json:"url,omitempty"`
}

// ProblemDetail defines model for ProblemDetail.
type ProblemDetail struct {
	// Content content of the problem
	Content   *string            `json:"content,omitempty"`
	CreatedAt *time.Time         `json:"createdAt,omitempty"`
	DomainId  openapi_types.UUID `json:"domainId"`

	// Hidden is the problem hidden
	Hidden         *bool               `json:"hidden,omitempty"`
	Id             uuid.UUID           `json:"id"`
	Languages      *[]string           `json:"languages,omitempty"`
	NumAccept      *int                `json:"numAccept,omitempty"`
	NumSubmit      *int                `json:"numSubmit,omitempty"`
	OwnerId        *openapi_types.UUID `json:"ownerId,omitempty"`
	ProblemGroupId *openapi_types.UUID `json:"problemGroupId,omitempty"`

	// Title title of the problem
	Title     string     `json:"title"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// Url (unique) url of the domain
	Url *string `json:"url,omitempty"`
}

// ProblemDetailWithLatestRecord defines model for ProblemDetailWithLatestRecord.
type ProblemDetailWithLatestRecord struct {
	// Content content of the problem
	Content   *string            `json:"content,omitempty"`
	CreatedAt *time.Time         `json:"createdAt,omitempty"`
	DomainId  openapi_types.UUID `json:"domainId"`

	// Hidden is the problem hidden
	Hidden         *bool               `json:"hidden,omitempty"`
	Id             uuid.UUID           `json:"id"`
	Languages      *[]string           `json:"languages,omitempty"`
	LatestRecord   *RecordPreview      `json:"latestRecord,omitempty"`
	NumAccept      *int                `json:"numAccept,omitempty"`
	NumSubmit      *int                `json:"numSubmit,omitempty"`
	OwnerId        *openapi_types.UUID `json:"ownerId,omitempty"`
	ProblemGroupId *openapi_types.UUID `json:"problemGroupId,omitempty"`

	// Title title of the problem
	Title     string     `json:"title"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// Url (unique) url of the domain
	Url *string `json:"url,omitempty"`
}

// ProblemEdit defines model for ProblemEdit.
type ProblemEdit struct {
	Content *string `json:"content,omitempty"`
	Hidden  *bool   `json:"hidden,omitempty"`
	Title   *string `json:"title,omitempty"`
	Url     *string `json:"url,omitempty"`
}

// ProblemGroup defines model for ProblemGroup.
type ProblemGroup struct {
	Id uuid.UUID `json:"id"`
}

// ProblemGroupList defines model for ProblemGroupList.
type ProblemGroupList struct {
	Count   *int            `json:"count,omitempty"`
	Results *[]ProblemGroup `json:"results,omitempty"`
}

// ProblemList defines model for ProblemList.
type ProblemList struct {
	Count   *int       `json:"count,omitempty"`
	Results *[]Problem `json:"results,omitempty"`
}

// ProblemPermission defines model for ProblemPermission.
type ProblemPermission struct {
	Create     *bool `json:"create,omitempty"`
	Edit       *bool `json:"edit,omitempty"`
	Submit     *bool `json:"submit,omitempty"`
	View       *bool `json:"view,omitempty"`
	ViewConfig *bool `json:"viewConfig,omitempty"`
	ViewHidden *bool `json:"viewHidden,omitempty"`
}

// ProblemPreviewWithLatestRecord defines model for ProblemPreviewWithLatestRecord.
type ProblemPreviewWithLatestRecord struct {
	// Hidden is the problem hidden
	Hidden       *bool               `json:"hidden,omitempty"`
	Id           uuid.UUID           `json:"id"`
	LatestRecord *RecordPreview      `json:"latestRecord,omitempty"`
	OwnerId      *openapi_types.UUID `json:"ownerId,omitempty"`

	// Title title of the problem
	Title string `json:"title"`

	// Url (unique) url of the domain
	Url *string `json:"url,omitempty"`
}

// ProblemSet defines model for ProblemSet.
type ProblemSet struct {
	// Content content of the problem set
	Content  *string            `json:"content,omitempty"`
	DomainId openapi_types.UUID `json:"domainId"`

	// DueAt the problem set is due at this date
	DueAt *time.Time `json:"dueAt,omitempty"`

	// Hidden whether the problem set is hidden
	Hidden *bool     `json:"hidden,omitempty"`
	Id     uuid.UUID `json:"id"`

	// LockAt the problem set is locked after this date
	LockAt    *time.Time          `json:"lockAt,omitempty"`
	NumAccept *int                `json:"numAccept,omitempty"`
	NumSubmit *int                `json:"numSubmit,omitempty"`
	OwnerId   *openapi_types.UUID `json:"ownerId,omitempty"`

	// ScoreboardHidden whether the scoreboard of the problem set is hidden
	ScoreboardHidden *bool `json:"scoreboardHidden,omitempty"`

	// Title title of the problem set
	Title string `json:"title"`

	// UnlockAt the problem set is unlocked after this date
	UnlockAt *time.Time `json:"unlockAt,omitempty"`

	// Url (unique) url of the domain
	Url *string `json:"url,omitempty"`
}

// ProblemSetAddProblem defines model for ProblemSetAddProblem.
type ProblemSetAddProblem struct {
	// Position the position of the problem in the problem set. if None, append to the end of the problems list.
	Position *int `json:"position,omitempty"`

	// Problem url or id of the problem
	Problem string `json:"problem"`
}

// ProblemSetCreate defines model for ProblemSetCreate.
type ProblemSetCreate struct {
	// Content content of the problem set
	Content *string    `json:"content,omitempty"`
	DueAt   *time.Time `json:"dueAt,omitempty"`

	// Hidden whether the problem set is hidden
	Hidden *bool      `json:"hidden,omitempty"`
	LockAt *time.Time `json:"lockAt,omitempty"`

	// ScoreboardHidden whether the scoreboard of the problem set is hidden
	ScoreboardHidden *bool `json:"scoreboardHidden,omitempty"`

	// Title title of the problem set
	Title    string     `json:"title"`
	UnlockAt *time.Time `json:"unlockAt,omitempty"`

	// Url (unique) url of the domain
	Url *string `json:"url,omitempty"`
}

// ProblemSetDetail defines model for ProblemSetDetail.
type ProblemSetDetail struct {
	// Content content of the problem set
	Content   *string            `json:"content,omitempty"`
	CreatedAt *time.Time         `json:"createdAt,omitempty"`
	DomainId  openapi_types.UUID `json:"domainId"`

	// DueAt the problem set is due at this date
	DueAt *time.Time `json:"dueAt,omitempty"`

	// Hidden whether the problem set is hidden
	Hidden *bool     `json:"hidden,omitempty"`
	Id     uuid.UUID `json:"id"`

	// LockAt the problem set is locked after this date
	LockAt    *time.Time                        `json:"lockAt,omitempty"`
	NumAccept *int                              `json:"numAccept,omitempty"`
	NumSubmit *int                              `json:"numSubmit,omitempty"`
	OwnerId   *openapi_types.UUID               `json:"ownerId,omitempty"`
	Problems  *[]ProblemPreviewWithLatestRecord `json:"problems,omitempty"`

	// ScoreboardHidden whether the scoreboard of the problem set is hidden
	ScoreboardHidden *bool `json:"scoreboardHidden,omitempty"`

	// Title title of the problem set
	Title string `json:"title"`

	// UnlockAt the problem set is unlocked after this date
	UnlockAt  *time.Time `json:"unlockAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// Url (unique) url of the domain
	Url *string `json:"url,omitempty"`
}

// ProblemSetEdit defines model for ProblemSetEdit.
type ProblemSetEdit struct {
	Content          *string    `json:"content,omitempty"`
	DueAt            *time.Time `json:"dueAt,omitempty"`
	Hidden           *bool      `json:"hidden,omitempty"`
	LockAt           *time.Time `json:"lockAt,omitempty"`
	ScoreboardHidden *bool      `json:"scoreboardHidden,omitempty"`
	Title            *string    `json:"title,omitempty"`
	UnlockAt         *time.Time `json:"unlockAt,omitempty"`
	Url              *string    `json:"url,omitempty"`
}

// ProblemSetList defines model for ProblemSetList.
type ProblemSetList struct {
	Count   *int          `json:"count,omitempty"`
	Results *[]ProblemSet `json:"results,omitempty"`
}

// ProblemSetPermission defines model for ProblemSetPermission.
type ProblemSetPermission struct {
	Claim      *bool `json:"claim,omitempty"`
	Create     *bool `json:"create,omitempty"`
	Edit       *bool `json:"edit,omitempty"`
	Manage     *bool `json:"manage,omitempty"`
	Scoreboard *bool `json:"scoreboard,omitempty"`
	View       *bool `json:"view,omitempty"`
	ViewConfig *bool `json:"viewConfig,omitempty"`
	ViewHidden *bool `json:"viewHidden,omitempty"`
}

// ProblemSetUpdateProblem defines model for ProblemSetUpdateProblem.
type ProblemSetUpdateProblem struct {
	// Position the position of the problem in the problem set. if None, append to the end of the problems list.
	Position *int `json:"position,omitempty"`
}

// ProblemSolutionSubmit defines model for ProblemSolutionSubmit.
type ProblemSolutionSubmit struct {
	Files    []openapi_types.File `json:"files"`
	Language string               `json:"language"`
}

// ProblemWithLatestRecord defines model for ProblemWithLatestRecord.
type ProblemWithLatestRecord struct {
	DomainId openapi_types.UUID `json:"domainId"`

	// Hidden is the problem hidden
	Hidden         *bool               `json:"hidden,omitempty"`
	Id             uuid.UUID           `json:"id"`
	LatestRecord   *RecordPreview      `json:"latestRecord,omitempty"`
	NumAccept      *int                `json:"numAccept,omitempty"`
	NumSubmit      *int                `json:"numSubmit,omitempty"`
	OwnerId        *openapi_types.UUID `json:"ownerId,omitempty"`
	ProblemGroupId *openapi_types.UUID `json:"problemGroupId,omitempty"`

	// Title title of the problem
	Title string `json:"title"`

	// Url (unique) url of the domain
	Url *string `json:"url,omitempty"`
}

// ProblemWithLatestRecordList defines model for ProblemWithLatestRecordList.
type ProblemWithLatestRecordList struct {
	Count   *int                       `json:"count,omitempty"`
	Results *[]ProblemWithLatestRecord `json:"results,omitempty"`
}

// ListProblemSetsParams defines parameters for ListProblemSets.
type ListProblemSetsParams struct{ Pagination }

// ListProblemsParams defines parameters for ListProblems.
type ListProblemsParams struct{ Pagination }

// ListProblemConfigCommitsParams defines parameters for ListProblemConfigCommits.
type ListProblemConfigCommitsParams struct{ Pagination }

// UpdateProblemConfigByArchiveParams defines parameters for UpdateProblemConfigByArchive.
type UpdateProblemConfigByArchiveParams struct {
	ConfigJsonOnMissing *ConfigMissing `form:"configJsonOnMissing,omitempty" json:"configJsonOnMissing,omitempty"`
}

// DiffProblemConfigDefaultBranchParams defines parameters for DiffProblemConfigDefaultBranch.
type DiffProblemConfigDefaultBranchParams struct {
	// After return items after this value
	After *string `form:"after,omitempty" json:"after,omitempty"`

	// Amount how many items to return
	Amount *int `form:"amount,omitempty" json:"amount,omitempty"`

	// Delimiter delimiter used to group common prefixes by
	Delimiter *string `form:"delimiter,omitempty" json:"delimiter,omitempty"`

	// Prefix return items prefixed with this value
	Prefix *string `form:"prefix,omitempty" json:"prefix,omitempty"`
}

// ListLatestProblemConfigObjectsUnderAGivenPrefixParams defines parameters for ListLatestProblemConfigObjectsUnderAGivenPrefix.
type ListLatestProblemConfigObjectsUnderAGivenPrefixParams struct {
	// After return items after this value
	After *string `form:"after,omitempty" json:"after,omitempty"`

	// Amount how many items to return
	Amount *int `form:"amount,omitempty" json:"amount,omitempty"`

	// Delimiter delimiter used to group common prefixes by
	Delimiter *string `form:"delimiter,omitempty" json:"delimiter,omitempty"`

	// Prefix return items prefixed with this value
	Prefix *string `form:"prefix,omitempty" json:"prefix,omitempty"`
}

// DownloadProblemConfigArchiveParams defines parameters for DownloadProblemConfigArchive.
type DownloadProblemConfigArchiveParams struct {
	ArchiveFormat *ArchiveType `form:"archiveFormat,omitempty" json:"archiveFormat,omitempty"`
}

// ListProblemGroupsParams defines parameters for ListProblemGroups.
type ListProblemGroupsParams struct{ Pagination }

// Diff defines model for Diff.
type Diff struct {
	Path string `json:"path"`

	// PathType An enumeration.
	PathType  PathTypeEnum `json:"pathType"`
	SizeBytes *int         `json:"sizeBytes,omitempty"`

	// Type An enumeration.
	Type DiffTypeEnum `json:"type"`
}

// DiffList defines model for DiffList.
type DiffList struct {
	Pagination Pagination `json:"pagination"`
	Results    []Diff     `json:"results"`
}

// FileUpload defines model for FileUpload.
type FileUpload struct {
	File openapi_types.File `json:"file"`
}

// ObjectStats defines model for ObjectStats.
type ObjectStats struct {
	Checksum    string  `json:"checksum"`
	ContentType *string `json:"contentType,omitempty"`
	Mtime       int     `json:"mtime"`
	Path        string  `json:"path"`

	// PathType An enumeration.
	PathType  PathTypeEnum `json:"pathType"`
	SizeBytes *int         `json:"sizeBytes,omitempty"`
}

// ObjectStatsList defines model for ObjectStatsList.
type ObjectStatsList struct {
	Pagination Pagination    `json:"pagination"`
	Results    []ObjectStats `json:"results"`
}
