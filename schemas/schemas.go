// Package schemas provides primitives to interact with the openapi HTTP API.
package schemas

import (
	"encoding/json"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

const (
	HTTPBearerScopes = "HTTPBearer.Scopes"
)

// Defines values for ArchiveType.
const (
	Rar     ArchiveType = "rar"
	Tar     ArchiveType = "tar"
	Unknown ArchiveType = "unknown"
	Zip     ArchiveType = "zip"
)

// Defines values for ConfigMissing.
const (
	RaiseError ConfigMissing = "raise_error"
	UseDefault ConfigMissing = "use_default"
	UseOld     ConfigMissing = "use_old"
)

// Defines values for DiffTypeEnum.
const (
	Added    DiffTypeEnum = "added"
	Changed  DiffTypeEnum = "changed"
	Conflict DiffTypeEnum = "conflict"
	Removed  DiffTypeEnum = "removed"
)

// Defines values for JWTAccessTokenCategory.
const (
	JWTAccessTokenCategoryOauth JWTAccessTokenCategory = "oauth"
	JWTAccessTokenCategoryUser  JWTAccessTokenCategory = "user"
)

// Defines values for PathTypeEnum.
const (
	CommonPrefix PathTypeEnum = "common_prefix"
	Object       PathTypeEnum = "object"
)

// Defines values for RecordCaseResult.
const (
	RecordCaseResultAccepted            RecordCaseResult = "accepted"
	RecordCaseResultCanceled            RecordCaseResult = "canceled"
	RecordCaseResultCompileError        RecordCaseResult = "compile_error"
	RecordCaseResultEtc                 RecordCaseResult = "etc"
	RecordCaseResultMemoryLimitExceeded RecordCaseResult = "memory_limit_exceeded"
	RecordCaseResultOutputLimitExceeded RecordCaseResult = "output_limit_exceeded"
	RecordCaseResultRuntimeError        RecordCaseResult = "runtime_error"
	RecordCaseResultSystemError         RecordCaseResult = "system_error"
	RecordCaseResultTimeLimitExceeded   RecordCaseResult = "time_limit_exceeded"
	RecordCaseResultWrongAnswer         RecordCaseResult = "wrong_answer"
)

// Defines values for RecordState.
const (
	RecordStateAccepted   RecordState = "accepted"
	RecordStateCompiling  RecordState = "compiling"
	RecordStateFailed     RecordState = "failed"
	RecordStateFetched    RecordState = "fetched"
	RecordStateJudging    RecordState = "judging"
	RecordStateProcessing RecordState = "processing"
	RecordStateQueueing   RecordState = "queueing"
	RecordStateRejected   RecordState = "rejected"
	RecordStateRetrying   RecordState = "retrying"
	RecordStateRunning    RecordState = "running"
)

// Defines values for LoginParamsResponseType.
const (
	LoginParamsResponseTypeJson     LoginParamsResponseType = "json"
	LoginParamsResponseTypeRedirect LoginParamsResponseType = "redirect"
)

// Defines values for LogoutParamsResponseType.
const (
	LogoutParamsResponseTypeJson     LogoutParamsResponseType = "json"
	LogoutParamsResponseTypeRedirect LogoutParamsResponseType = "redirect"
)

// Defines values for OauthAuthorizeParamsResponseType.
const (
	OauthAuthorizeParamsResponseTypeJson     OauthAuthorizeParamsResponseType = "json"
	OauthAuthorizeParamsResponseTypeRedirect OauthAuthorizeParamsResponseType = "redirect"
)

// Defines values for RefreshParamsResponseType.
const (
	RefreshParamsResponseTypeJson     RefreshParamsResponseType = "json"
	RefreshParamsResponseTypeRedirect RefreshParamsResponseType = "redirect"
)

// Defines values for RegisterParamsResponseType.
const (
	RegisterParamsResponseTypeJson     RegisterParamsResponseType = "json"
	RegisterParamsResponseTypeRedirect RegisterParamsResponseType = "redirect"
)

// Defines values for GetTokenParamsResponseType.
const (
	GetTokenParamsResponseTypeJson     GetTokenParamsResponseType = "json"
	GetTokenParamsResponseTypeRedirect GetTokenParamsResponseType = "redirect"
)

// ArchiveType An enumeration.
type ArchiveType string

// AuthTokens defines model for AuthTokens.
type AuthTokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	TokenType    string `json:"tokenType"`
}

// AuthTokensResp defines model for AuthTokensResp.
type AuthTokensResp struct {
	Data *AuthTokens `json:"data,omitempty"`
	BizError
}

// AuthTokensWithLakefs defines model for AuthTokensWithLakefs.
type AuthTokensWithLakefs struct {
	AccessKeyId     string `json:"accessKeyId"`
	AccessToken     string `json:"accessToken"`
	RefreshToken    string `json:"refreshToken"`
	SecretAccessKey string `json:"secretAccessKey"`
	TokenType       string `json:"tokenType"`
}

// AuthTokensWithLakefsResp defines model for AuthTokensWithLakefsResp.
type AuthTokensWithLakefsResp struct {
	Data *AuthTokensWithLakefs `json:"data,omitempty"`
	BizError
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

// ConfigMissing An enumeration.
type ConfigMissing string

// Detail defines model for Detail.
type Detail struct {
	Detail string `json:"detail"`
}

// Diff defines model for Diff.
type Diff struct {
	Path string `json:"path"`

	// PathType An enumeration.
	PathType  PathTypeEnum `json:"path_type"`
	SizeBytes *int         `json:"size_bytes,omitempty"`

	// Type An enumeration.
	Type DiffTypeEnum `json:"type"`
}

// DiffList defines model for DiffList.
type DiffList struct {
	Pagination Pagination `json:"pagination"`
	Results    []Diff     `json:"results"`
}

// DiffListResp defines model for DiffListResp.
type DiffListResp struct {
	Data *DiffList `json:"data,omitempty"`
	BizError
}

// DiffTypeEnum An enumeration.
type DiffTypeEnum string

// Domain defines model for Domain.
type Domain struct {
	// Bulletin bulletin of the domain
	Bulletin *string `json:"bulletin,omitempty"`

	// Gravatar gravatar url of the domain
	Gravatar *string `json:"gravatar,omitempty"`

	// Group group name of the domain
	Group *string `json:"group,omitempty"`

	// Hidden is the domain hidden
	Hidden *bool              `json:"hidden,omitempty"`
	Id     openapi_types.UUID `json:"id"`

	// Name displayed name of the domain
	Name    string              `json:"name"`
	OwnerId *openapi_types.UUID `json:"ownerId,omitempty"`

	// Url (unique) url of the domain
	Url *string `json:"url,omitempty"`
}

// DomainCreate defines model for DomainCreate.
type DomainCreate struct {
	// Bulletin bulletin of the domain
	Bulletin *string `json:"bulletin,omitempty"`

	// Gravatar gravatar url of the domain
	Gravatar *string `json:"gravatar,omitempty"`

	// Group group name of the domain
	Group *string `json:"group,omitempty"`

	// Hidden is the domain hidden
	Hidden *bool `json:"hidden,omitempty"`

	// Name displayed name of the domain
	Name string `json:"name"`

	// Url (unique) url of the domain
	Url *string `json:"url,omitempty" validate:"domain_url"`
}

// DomainDetail defines model for DomainDetail.
type DomainDetail struct {
	// Bulletin bulletin of the domain
	Bulletin  *string    `json:"bulletin,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// Gravatar gravatar url of the domain
	Gravatar *string `json:"gravatar,omitempty"`

	// Group group name of the domain
	Group *string `json:"group,omitempty"`

	// Hidden is the domain hidden
	Hidden *bool              `json:"hidden,omitempty"`
	Id     openapi_types.UUID `json:"id"`

	// Name displayed name of the domain
	Name      string              `json:"name"`
	OwnerId   *openapi_types.UUID `json:"ownerId,omitempty"`
	Tag       *string             `json:"tag,omitempty"`
	UpdatedAt *time.Time          `json:"updatedAt,omitempty"`

	// Url (unique) url of the domain
	Url *string `json:"url,omitempty"`
}

// DomainDetailResp defines model for DomainDetailResp.
type DomainDetailResp struct {
	Data *DomainDetail `json:"data,omitempty"`
	BizError
}

// DomainEdit defines model for DomainEdit.
type DomainEdit struct {
	Bulletin *string `json:"bulletin,omitempty"`
	Gravatar *string `json:"gravatar,omitempty"`
	Group    *string `json:"group,omitempty"`
	Hidden   *bool   `json:"hidden,omitempty"`
	Name     *string `json:"name,omitempty"`
	Url      *string `json:"url,omitempty"`
}

// DomainInvitation defines model for DomainInvitation.
type DomainInvitation struct {
	// Code invitation code
	Code     string             `json:"code"`
	DomainId openapi_types.UUID `json:"domainId"`

	// ExpireAt expire time of invitation
	ExpireAt *time.Time         `json:"expireAt,omitempty"`
	Id       openapi_types.UUID `json:"id"`

	// Role domain role after invitation
	Role *string `json:"role,omitempty"`

	// Url (unique) url of the domain
	Url *string `json:"url,omitempty"`
}

// DomainInvitationCreate defines model for DomainInvitationCreate.
type DomainInvitationCreate struct {
	// Code invitation code
	Code     string     `json:"code"`
	ExpireAt *time.Time `json:"expireAt,omitempty"`

	// Role domain role after invitation
	Role *string `json:"role,omitempty"`

	// Url (unique) url of the domain
	Url *string `json:"url,omitempty"`
}

// DomainInvitationEdit defines model for DomainInvitationEdit.
type DomainInvitationEdit struct {
	// Code invitation code
	Code *string `json:"code,omitempty"`

	// ExpireAt expire time of invitation
	ExpireAt *time.Time `json:"expireAt,omitempty"`

	// Role domain role after invitation
	Role *string `json:"role,omitempty"`
}

// DomainInvitationList defines model for DomainInvitationList.
type DomainInvitationList struct {
	Count   *int                `json:"count,omitempty"`
	Results *[]DomainInvitation `json:"results,omitempty"`
}

// DomainInvitationListResp defines model for DomainInvitationListResp.
type DomainInvitationListResp struct {
	Data *DomainInvitationList `json:"data,omitempty"`
	BizError
}

// DomainInvitationResp defines model for DomainInvitationResp.
type DomainInvitationResp struct {
	Data *DomainInvitation `json:"data,omitempty"`
	BizError
}

// DomainList defines model for DomainList.
type DomainList struct {
	Count   *int      `json:"count,omitempty"`
	Results *[]Domain `json:"results,omitempty"`
}

// DomainListResp defines model for DomainListResp.
type DomainListResp struct {
	Data *DomainList `json:"data,omitempty"`
	BizError
}

// DomainPermission All permissions in a domain
type DomainPermission struct {
	General    *GeneralPermission    `json:"general,omitempty"`
	Problem    *ProblemPermission    `json:"problem,omitempty"`
	ProblemSet *ProblemSetPermission `json:"problemSet,omitempty"`
	Record     *RecordPermission     `json:"record,omitempty"`
}

// DomainResp defines model for DomainResp.
type DomainResp struct {
	Data *Domain `json:"data,omitempty"`
	BizError
}

// DomainRole defines model for DomainRole.
type DomainRole struct {
	DomainId openapi_types.UUID `json:"domainId"`
	Id       openapi_types.UUID `json:"id"`

	// Permission All permissions in a domain
	Permission DomainPermission `json:"permission"`
	Role       string           `json:"role"`
}

// DomainRoleCreate defines model for DomainRoleCreate.
type DomainRoleCreate struct {
	// Permission All permissions in a domain
	Permission DomainPermission `json:"permission"`
	Role       string           `json:"role"`
}

// DomainRoleDetail defines model for DomainRoleDetail.
type DomainRoleDetail struct {
	CreatedAt *time.Time         `json:"createdAt,omitempty"`
	DomainId  openapi_types.UUID `json:"domainId"`
	Id        openapi_types.UUID `json:"id"`

	// Permission All permissions in a domain
	Permission DomainPermission `json:"permission"`
	Role       string           `json:"role"`
	UpdatedAt  *time.Time       `json:"updatedAt,omitempty"`
}

// DomainRoleDetailResp defines model for DomainRoleDetailResp.
type DomainRoleDetailResp struct {
	Data *DomainRoleDetail `json:"data,omitempty"`
	BizError
}

// DomainRoleEdit defines model for DomainRoleEdit.
type DomainRoleEdit struct {
	// Permission The permission which needs to be updated
	Permission *map[string]any `json:"permission,omitempty"`
}

// DomainRoleList defines model for DomainRoleList.
type DomainRoleList struct {
	Count   *int          `json:"count,omitempty"`
	Results *[]DomainRole `json:"results,omitempty"`
}

// DomainRoleListResp defines model for DomainRoleListResp.
type DomainRoleListResp struct {
	Data *DomainRoleList `json:"data,omitempty"`
	BizError
}

// DomainRoleResp defines model for DomainRoleResp.
type DomainRoleResp struct {
	Data *DomainRole `json:"data,omitempty"`
	BizError
}

// DomainTag defines model for DomainTag.
type DomainTag = string

// DomainTagList defines model for DomainTagList.
type DomainTagList struct {
	Count   *int         `json:"count,omitempty"`
	Results *[]DomainTag `json:"results,omitempty"`
}

// DomainTagListResp defines model for DomainTagListResp.
type DomainTagListResp struct {
	Data *DomainTagList `json:"data,omitempty"`
	BizError
}

// DomainTransfer defines model for DomainTransfer.
type DomainTransfer struct {
	// TargetUser 'me' or id of the user
	TargetUser string `json:"targetUser"`
}

// DomainUserAdd defines model for DomainUserAdd.
type DomainUserAdd struct {
	Role *string `json:"role,omitempty"`

	// User 'me' or id of the user
	User string `json:"user"`
}

// DomainUserPermission defines model for DomainUserPermission.
type DomainUserPermission struct {
	DomainId openapi_types.UUID `json:"domainId"`

	// Permission All permissions in a domain
	Permission DomainPermission   `json:"permission"`
	Role       string             `json:"role"`
	UserId     openapi_types.UUID `json:"userId"`
}

// DomainUserPermissionResp defines model for DomainUserPermissionResp.
type DomainUserPermissionResp struct {
	Data *DomainUserPermission `json:"data,omitempty"`
	BizError
}

// DomainUserUpdate defines model for DomainUserUpdate.
type DomainUserUpdate struct {
	Role *string `json:"role,omitempty"`
}

// FileUpload defines model for FileUpload.
type FileUpload struct {
	File openapi_types.File `json:"file"`
}

// GeneralPermission defines model for GeneralPermission.
type GeneralPermission struct {
	Edit           *bool `json:"edit,omitempty"`
	EditPermission *bool `json:"editPermission,omitempty"`
	UnlimitedQuota *bool `json:"unlimitedQuota,omitempty"`
	View           *bool `json:"view,omitempty"`
	ViewModBadge   *bool `json:"viewModBadge,omitempty"`
}

// HTTPValidationError defines model for HTTPValidationError.
type HTTPValidationError struct {
	Detail *[]ValidationError `json:"detail,omitempty"`
}

// JWTAccessToken defines model for JWTAccessToken.
type JWTAccessToken struct {
	Category  JWTAccessTokenCategory `json:"category"`
	Csrf      *string                `json:"csrf,omitempty"`
	Exp       int                    `json:"exp"`
	Fresh     *bool                  `json:"fresh,omitempty"`
	Gravatar  *string                `json:"gravatar,omitempty"`
	Iat       int                    `json:"iat"`
	IsActive  bool                   `json:"isActive"`
	Jti       string                 `json:"jti"`
	Nbf       int                    `json:"nbf"`
	OauthName *string                `json:"oauthName,omitempty"`
	Role      *string                `json:"role,omitempty"`
	Sub       string                 `json:"sub"`
	Type      string                 `json:"type"`
	Username  string                 `json:"username"`
}

// JWTAccessTokenCategory defines model for JWTAccessToken.Category.
type JWTAccessTokenCategory string

// JWTAccessTokenResp defines model for JWTAccessTokenResp.
type JWTAccessTokenResp struct {
	Data *JWTAccessToken `json:"data,omitempty"`
	BizError
}

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

// JudgerCredentialsResp defines model for JudgerCredentialsResp.
type JudgerCredentialsResp struct {
	Data *JudgerCredentials `json:"data,omitempty"`
	BizError
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

// JudgerDetailListResp defines model for JudgerDetailListResp.
type JudgerDetailListResp struct {
	Data *JudgerDetailList `json:"data,omitempty"`
	BizError
}

// Language defines model for Language.
type Language struct {
	CaseDefault  *Case     `json:"caseDefault,omitempty"`
	Cases        *[]Case   `json:"cases,omitempty"`
	CompileArgs  *[]string `json:"compileArgs,omitempty"`
	CompileFiles *[]string `json:"compileFiles,omitempty"`
	Name         string    `json:"name"`
}

// LanguageDefault defines model for LanguageDefault.
type LanguageDefault struct {
	CaseDefault  *Case     `json:"caseDefault,omitempty"`
	Cases        *[]Case   `json:"cases,omitempty"`
	CompileArgs  *[]string `json:"compileArgs,omitempty"`
	CompileFiles *[]string `json:"compileFiles,omitempty"`
}

// OAuth2Client defines model for OAuth2Client.
type OAuth2Client struct {
	DisplayName string `json:"displayName"`
	Icon        string `json:"icon"`
	OauthName   string `json:"oauthName"`
}

// OAuth2ClientList defines model for OAuth2ClientList.
type OAuth2ClientList struct {
	Count   *int            `json:"count,omitempty"`
	Results *[]OAuth2Client `json:"results,omitempty"`
}

// OAuth2ClientListResp defines model for OAuth2ClientListResp.
type OAuth2ClientListResp struct {
	Data *OAuth2ClientList `json:"data,omitempty"`
	BizError
}

// OAuth2PasswordRequestForm defines model for OAuth2PasswordRequestForm.
type OAuth2PasswordRequestForm struct {
	ClientId     *string `json:"clientId,omitempty"`
	ClientSecret *string `json:"clientSecret,omitempty"`
	GrantType    *string `json:"grantType,omitempty"`
	Password     *string `json:"password,omitempty"`
	Scope        *string `json:"scope,omitempty"`
	Username     *string `json:"username,omitempty"`
}

// ObjectStats defines model for ObjectStats.
type ObjectStats struct {
	Checksum    string  `json:"checksum"`
	ContentType *string `json:"content_type,omitempty"`
	Mtime       int     `json:"mtime"`
	Path        string  `json:"path"`

	// PathType An enumeration.
	PathType  PathTypeEnum `json:"path_type"`
	SizeBytes *int         `json:"size_bytes,omitempty"`
}

// ObjectStatsList defines model for ObjectStatsList.
type ObjectStatsList struct {
	Pagination Pagination    `json:"pagination"`
	Results    []ObjectStats `json:"results"`
}

// ObjectStatsListResp defines model for ObjectStatsListResp.
type ObjectStatsListResp struct {
	Data *ObjectStatsList `json:"data,omitempty"`
	BizError
}

// Pagination defines model for Pagination.
type Pagination struct {
	HasMore    bool   `json:"has_more"`
	MaxPerPage int    `json:"max_per_page"`
	NextOffset string `json:"next_offset"`
	Results    int    `json:"results"`
}

// PathTypeEnum An enumeration.
type PathTypeEnum string

// Problem defines model for Problem.
type Problem struct {
	DomainId openapi_types.UUID `json:"domainId"`

	// Hidden is the problem hidden
	Hidden         *bool               `json:"hidden,omitempty"`
	Id             openapi_types.UUID  `json:"id"`
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
	Id            openapi_types.UUID  `json:"id"`
	UpdatedAt     *time.Time          `json:"updatedAt,omitempty"`
}

// ProblemConfigDataDetailResp defines model for ProblemConfigDataDetailResp.
type ProblemConfigDataDetailResp struct {
	Data *ProblemConfigDataDetail `json:"data,omitempty"`
	BizError
}

// ProblemConfigDetail defines model for ProblemConfigDetail.
type ProblemConfigDetail struct {
	CommitId      *string             `json:"commitId,omitempty"`
	CommitMessage *string             `json:"commitMessage,omitempty"`
	CommitterId   *openapi_types.UUID `json:"committerId,omitempty"`
	CreatedAt     *time.Time          `json:"createdAt,omitempty"`
	DataVersion   *int                `json:"dataVersion,omitempty"`
	Id            openapi_types.UUID  `json:"id"`
	UpdatedAt     *time.Time          `json:"updatedAt,omitempty"`
}

// ProblemConfigDetailList defines model for ProblemConfigDetailList.
type ProblemConfigDetailList struct {
	Count   *int                   `json:"count,omitempty"`
	Results *[]ProblemConfigDetail `json:"results,omitempty"`
}

// ProblemConfigDetailListResp defines model for ProblemConfigDetailListResp.
type ProblemConfigDetailListResp struct {
	Data *ProblemConfigDetailList `json:"data,omitempty"`
	BizError
}

// ProblemConfigDetailResp defines model for ProblemConfigDetailResp.
type ProblemConfigDetailResp struct {
	Data *ProblemConfigDetail `json:"data,omitempty"`
	BizError
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
	Id             openapi_types.UUID  `json:"id"`
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

// ProblemDetailResp defines model for ProblemDetailResp.
type ProblemDetailResp struct {
	Data *ProblemDetail `json:"data,omitempty"`
	BizError
}

// ProblemDetailWithLatestRecord defines model for ProblemDetailWithLatestRecord.
type ProblemDetailWithLatestRecord struct {
	// Content content of the problem
	Content   *string            `json:"content,omitempty"`
	CreatedAt *time.Time         `json:"createdAt,omitempty"`
	DomainId  openapi_types.UUID `json:"domainId"`

	// Hidden is the problem hidden
	Hidden         *bool               `json:"hidden,omitempty"`
	Id             openapi_types.UUID  `json:"id"`
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

// ProblemDetailWithLatestRecordResp defines model for ProblemDetailWithLatestRecordResp.
type ProblemDetailWithLatestRecordResp struct {
	Data *ProblemDetailWithLatestRecord `json:"data,omitempty"`
	BizError
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
	Id openapi_types.UUID `json:"id"`
}

// ProblemGroupList defines model for ProblemGroupList.
type ProblemGroupList struct {
	Count   *int            `json:"count,omitempty"`
	Results *[]ProblemGroup `json:"results,omitempty"`
}

// ProblemGroupListResp defines model for ProblemGroupListResp.
type ProblemGroupListResp struct {
	Data *ProblemGroupList `json:"data,omitempty"`
	BizError
}

// ProblemList defines model for ProblemList.
type ProblemList struct {
	Count   *int       `json:"count,omitempty"`
	Results *[]Problem `json:"results,omitempty"`
}

// ProblemListResp defines model for ProblemListResp.
type ProblemListResp struct {
	Data *ProblemList `json:"data,omitempty"`
	BizError
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
	Id           openapi_types.UUID  `json:"id"`
	LatestRecord *RecordPreview      `json:"latestRecord,omitempty"`
	OwnerId      *openapi_types.UUID `json:"ownerId,omitempty"`

	// Title title of the problem
	Title string `json:"title"`

	// Url (unique) url of the domain
	Url *string `json:"url,omitempty"`
}

// ProblemResp defines model for ProblemResp.
type ProblemResp struct {
	Data *Problem `json:"data,omitempty"`
	BizError
}

// ProblemSet defines model for ProblemSet.
type ProblemSet struct {
	// Content content of the problem set
	Content  *string            `json:"content,omitempty"`
	DomainId openapi_types.UUID `json:"domainId"`

	// DueAt the problem set is due at this date
	DueAt *time.Time `json:"dueAt,omitempty"`

	// Hidden whether the problem set is hidden
	Hidden *bool              `json:"hidden,omitempty"`
	Id     openapi_types.UUID `json:"id"`

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
	Hidden *bool              `json:"hidden,omitempty"`
	Id     openapi_types.UUID `json:"id"`

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

// ProblemSetDetailResp defines model for ProblemSetDetailResp.
type ProblemSetDetailResp struct {
	Data *ProblemSetDetail `json:"data,omitempty"`
	BizError
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

// ProblemSetListResp defines model for ProblemSetListResp.
type ProblemSetListResp struct {
	Data *ProblemSetList `json:"data,omitempty"`
	BizError
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

// ProblemSetResp defines model for ProblemSetResp.
type ProblemSetResp struct {
	Data *ProblemSet `json:"data,omitempty"`
	BizError
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
	Id             openapi_types.UUID  `json:"id"`
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

// ProblemWithLatestRecordListResp defines model for ProblemWithLatestRecordListResp.
type ProblemWithLatestRecordListResp struct {
	Data *ProblemWithLatestRecordList `json:"data,omitempty"`
	BizError
}

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

// RecordCaseResult An enumeration.
type RecordCaseResult string

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

// RecordDetailResp defines model for RecordDetailResp.
type RecordDetailResp struct {
	Data *RecordDetail `json:"data,omitempty"`
	BizError
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

// RecordListDetailListResp defines model for RecordListDetailListResp.
type RecordListDetailListResp struct {
	Data *RecordListDetailList `json:"data,omitempty"`
	BizError
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

// RecordResp defines model for RecordResp.
type RecordResp struct {
	Data *Record `json:"data,omitempty"`
	BizError
}

// RecordState An enumeration.
type RecordState string

// RecordSubmit defines model for RecordSubmit.
type RecordSubmit struct {
	JudgedAt *time.Time `json:"judgedAt,omitempty"`
	MemoryKb *int       `json:"memoryKb,omitempty"`
	Score    *int       `json:"score,omitempty"`

	// State An enumeration.
	State  *RecordState `json:"state,omitempty"`
	TimeMs *int         `json:"timeMs,omitempty"`
}

// Redirect defines model for Redirect.
type Redirect struct {
	RedirectUrl string `json:"redirectUrl"`
}

// RedirectResp defines model for RedirectResp.
type RedirectResp struct {
	Data *Redirect `json:"data,omitempty"`
	BizError
}

// User defines model for User.
type User struct {
	Gravatar *string            `json:"gravatar,omitempty"`
	Id       openapi_types.UUID `json:"id"`
	IsActive *bool              `json:"isActive,omitempty"`
	Role     *string            `json:"role,omitempty"`
	Username string             `json:"username"`
}

// UserCreate defines model for UserCreate.
type UserCreate struct {
	Email          *string `json:"email,omitempty"`
	OauthAccountId *string `json:"oauth_account_id,omitempty"`
	OauthName      *string `json:"oauth_name,omitempty"`
	Password       *string `json:"password,omitempty"`
	Username       *string `json:"username,omitempty"`
}

// UserDetail defines model for UserDetail.
type UserDetail struct {
	CreatedAt  *time.Time          `json:"createdAt,omitempty"`
	Email      openapi_types.Email `json:"email"`
	Gravatar   *string             `json:"gravatar,omitempty"`
	Id         openapi_types.UUID  `json:"id"`
	IsActive   *bool               `json:"isActive,omitempty"`
	LoginAt    time.Time           `json:"loginAt"`
	LoginIp    string              `json:"loginIp"`
	RealName   *string             `json:"realName,omitempty"`
	RegisterIp string              `json:"registerIp"`
	Role       *string             `json:"role,omitempty"`
	StudentId  *string             `json:"studentId,omitempty"`
	UpdatedAt  *time.Time          `json:"updatedAt,omitempty"`
	Username   string              `json:"username"`
}

// UserDetailResp defines model for UserDetailResp.
type UserDetailResp struct {
	Data *UserDetail `json:"data,omitempty"`
	BizError
}

// UserDetailWithDomainRole defines model for UserDetailWithDomainRole.
type UserDetailWithDomainRole struct {
	CreatedAt  *time.Time          `json:"createdAt,omitempty"`
	DomainId   *openapi_types.UUID `json:"domainId,omitempty"`
	DomainRole *string             `json:"domainRole,omitempty"`
	Email      openapi_types.Email `json:"email"`
	Gravatar   *string             `json:"gravatar,omitempty"`
	Id         openapi_types.UUID  `json:"id"`
	IsActive   *bool               `json:"isActive,omitempty"`
	LoginAt    time.Time           `json:"loginAt"`
	LoginIp    string              `json:"loginIp"`
	RealName   *string             `json:"realName,omitempty"`
	RegisterIp string              `json:"registerIp"`
	Role       *string             `json:"role,omitempty"`
	StudentId  *string             `json:"studentId,omitempty"`
	UpdatedAt  *time.Time          `json:"updatedAt,omitempty"`
	Username   string              `json:"username"`
}

// UserDetailWithDomainRoleList defines model for UserDetailWithDomainRoleList.
type UserDetailWithDomainRoleList struct {
	Count   *int                        `json:"count,omitempty"`
	Results *[]UserDetailWithDomainRole `json:"results,omitempty"`
}

// UserDetailWithDomainRoleListResp defines model for UserDetailWithDomainRoleListResp.
type UserDetailWithDomainRoleListResp struct {
	Data *UserDetailWithDomainRoleList `json:"data,omitempty"`
	BizError
}

// UserDetailWithDomainRoleResp defines model for UserDetailWithDomainRoleResp.
type UserDetailWithDomainRoleResp struct {
	Data *UserDetailWithDomainRole `json:"data,omitempty"`
	BizError
}

// UserEdit defines model for UserEdit.
type UserEdit struct {
	Gravatar *string `json:"gravatar,omitempty"`
}

// UserList defines model for UserList.
type UserList struct {
	Count   *int    `json:"count,omitempty"`
	Results *[]User `json:"results,omitempty"`
}

// UserListResp defines model for UserListResp.
type UserListResp struct {
	Data *UserList `json:"data,omitempty"`
	BizError
}

// UserPreview defines model for UserPreview.
type UserPreview struct {
	Gravatar *string            `json:"gravatar,omitempty"`
	Id       openapi_types.UUID `json:"id"`
	Username string             `json:"username"`
}

// UserPreviewResp defines model for UserPreviewResp.
type UserPreviewResp struct {
	Data *UserPreview `json:"data,omitempty"`
	BizError
}

// UserResetPassword defines model for UserResetPassword.
type UserResetPassword struct {
	CurrentPassword *string `json:"currentPassword,omitempty"`
	NewPassword     string  `json:"newPassword"`
}

// UserWithDomainRole defines model for UserWithDomainRole.
type UserWithDomainRole struct {
	DomainId   *openapi_types.UUID `json:"domainId,omitempty"`
	DomainRole *string             `json:"domainRole,omitempty"`
	Gravatar   *string             `json:"gravatar,omitempty"`
	Id         openapi_types.UUID  `json:"id"`
	Username   string              `json:"username"`
}

// UserWithDomainRoleList defines model for UserWithDomainRoleList.
type UserWithDomainRoleList struct {
	Count   *int                  `json:"count,omitempty"`
	Results *[]UserWithDomainRole `json:"results,omitempty"`
}

// UserWithDomainRoleListResp defines model for UserWithDomainRoleListResp.
type UserWithDomainRoleListResp struct {
	Data *UserWithDomainRoleList `json:"data,omitempty"`
	BizError
}

// UserWithDomainRoleResp defines model for UserWithDomainRoleResp.
type UserWithDomainRoleResp struct {
	Data *UserWithDomainRole `json:"data,omitempty"`
	BizError
}

// ValidationError defines model for ValidationError.
type ValidationError struct {
	Loc  []ValidationError_Loc_Item `json:"loc"`
	Msg  string                     `json:"msg"`
	Type string                     `json:"type"`
}

// ValidationErrorLoc0 defines model for .
type ValidationErrorLoc0 = string

// ValidationErrorLoc1 defines model for .
type ValidationErrorLoc1 = int

// ValidationError_Loc_Item defines model for ValidationError.loc.Item.
type ValidationError_Loc_Item struct {
	union json.RawMessage
}

// Version defines model for Version.
type Version struct {
	Git     string `json:"git"`
	Version string `json:"version"`
}

// AdminListDomainRolesParams defines parameters for AdminListDomainRoles.
type AdminListDomainRolesParams struct {
	// Ordering Comma separated list of ordering the results.
	// You may specify reverse orderings by prefixing the field name with '-'.
	//
	// Available fields: created_at,updated_at
	Ordering *string `form:"ordering,omitempty" json:"ordering,omitempty"`
	Offset   *int    `form:"offset,omitempty"   json:"offset,omitempty"`
	Limit    *int    `form:"limit,omitempty"    json:"limit,omitempty"`
}

// AdminListJudgersParams defines parameters for AdminListJudgers.
type AdminListJudgersParams struct {
	// Ordering Comma separated list of ordering the results.
	// You may specify reverse orderings by prefixing the field name with '-'.
	//
	// Available fields: created_at,updated_at
	Ordering *string `form:"ordering,omitempty" json:"ordering,omitempty"`
	Offset   *int    `form:"offset,omitempty"   json:"offset,omitempty"`
	Limit    *int    `form:"limit,omitempty"    json:"limit,omitempty"`
}

// AdminListUsersParams defines parameters for AdminListUsers.
type AdminListUsersParams struct {
	// Ordering Comma separated list of ordering the results.
	// You may specify reverse orderings by prefixing the field name with '-'.
	//
	// Available fields: created_at,updated_at
	Ordering *string `form:"ordering,omitempty" json:"ordering,omitempty"`
	Offset   *int    `form:"offset,omitempty"   json:"offset,omitempty"`
	Limit    *int    `form:"limit,omitempty"    json:"limit,omitempty"`
}

// AdminListUserDomainsParams defines parameters for AdminListUserDomains.
type AdminListUserDomainsParams struct {
	Role   *[]string `form:"role,omitempty"   json:"role,omitempty"`
	Groups *[]string `form:"groups,omitempty" json:"groups,omitempty"`

	// Ordering Comma separated list of ordering the results.
	// You may specify reverse orderings by prefixing the field name with '-'.
	//
	// Available fields: created_at,updated_at
	Ordering *string `form:"ordering,omitempty" json:"ordering,omitempty"`
	Offset   *int    `form:"offset,omitempty"   json:"offset,omitempty"`
	Limit    *int    `form:"limit,omitempty"    json:"limit,omitempty"`
}

// LoginParams defines parameters for Login.
type LoginParams struct {
	// Cookie Add Set/Delete-Cookie on response header
	Cookie       *bool                   `form:"cookie,omitempty" json:"cookie,omitempty"`
	ResponseType LoginParamsResponseType `form:"responseType"     json:"responseType"`

	// RedirectUrl The redirect url after the operation
	RedirectUrl *string `form:"redirectUrl,omitempty" json:"redirectUrl,omitempty"`
}

// LoginParamsResponseType defines parameters for Login.
type LoginParamsResponseType string

// LogoutParams defines parameters for Logout.
type LogoutParams struct {
	// Cookie Add Set/Delete-Cookie on response header
	Cookie       *bool                    `form:"cookie,omitempty" json:"cookie,omitempty"`
	ResponseType LogoutParamsResponseType `form:"responseType"     json:"responseType"`

	// RedirectUrl The redirect url after the operation
	RedirectUrl *string `form:"redirectUrl,omitempty" json:"redirectUrl,omitempty"`
}

// LogoutParamsResponseType defines parameters for Logout.
type LogoutParamsResponseType string

// OauthAuthorizeParams defines parameters for OauthAuthorize.
type OauthAuthorizeParams struct {
	Scopes *[]string `form:"scopes,omitempty" json:"scopes,omitempty"`

	// Cookie Add Set/Delete-Cookie on response header
	Cookie       *bool                            `form:"cookie,omitempty" json:"cookie,omitempty"`
	ResponseType OauthAuthorizeParamsResponseType `form:"responseType"     json:"responseType"`

	// RedirectUrl The redirect url after the operation
	RedirectUrl *string `form:"redirectUrl,omitempty" json:"redirectUrl,omitempty"`
}

// OauthAuthorizeParamsResponseType defines parameters for OauthAuthorize.
type OauthAuthorizeParamsResponseType string

// RefreshParams defines parameters for Refresh.
type RefreshParams struct {
	// Cookie Add Set/Delete-Cookie on response header
	Cookie       *bool                     `form:"cookie,omitempty" json:"cookie,omitempty"`
	ResponseType RefreshParamsResponseType `form:"responseType"     json:"responseType"`

	// RedirectUrl The redirect url after the operation
	RedirectUrl *string `form:"redirectUrl,omitempty" json:"redirectUrl,omitempty"`
}

// RefreshParamsResponseType defines parameters for Refresh.
type RefreshParamsResponseType string

// RegisterParams defines parameters for Register.
type RegisterParams struct {
	// Cookie Add Set/Delete-Cookie on response header
	Cookie       *bool                      `form:"cookie,omitempty" json:"cookie,omitempty"`
	ResponseType RegisterParamsResponseType `form:"responseType"     json:"responseType"`

	// RedirectUrl The redirect url after the operation
	RedirectUrl *string `form:"redirectUrl,omitempty" json:"redirectUrl,omitempty"`
}

// RegisterParamsResponseType defines parameters for Register.
type RegisterParamsResponseType string

// GetTokenParams defines parameters for GetToken.
type GetTokenParams struct {
	// Cookie Add Set/Delete-Cookie on response header
	Cookie       *bool                      `form:"cookie,omitempty" json:"cookie,omitempty"`
	ResponseType GetTokenParamsResponseType `form:"responseType"     json:"responseType"`

	// RedirectUrl The redirect url after the operation
	RedirectUrl *string `form:"redirectUrl,omitempty" json:"redirectUrl,omitempty"`
}

// GetTokenParamsResponseType defines parameters for GetToken.
type GetTokenParamsResponseType string

// ListDomainsParams defines parameters for ListDomains.
type ListDomainsParams struct {
	Roles  *[]string `form:"roles,omitempty"  json:"roles,omitempty"`
	Groups *[]string `form:"groups,omitempty" json:"groups,omitempty"`

	// Ordering Comma separated list of ordering the results.
	// You may specify reverse orderings by prefixing the field name with '-'.
	//
	// Available fields: created_at,updated_at
	Ordering *string `form:"ordering,omitempty" json:"ordering,omitempty"`
	Offset   *int    `form:"offset,omitempty"   json:"offset,omitempty"`
	Limit    *int    `form:"limit,omitempty"    json:"limit,omitempty"`
}

// SearchDomainGroupsParams defines parameters for SearchDomainGroups.
type SearchDomainGroupsParams struct {
	// Query search query
	Query string `form:"query" json:"query"`
}

// SearchDomainCandidatesParams defines parameters for SearchDomainCandidates.
type SearchDomainCandidatesParams struct {
	// Query search query
	Query string `form:"query" json:"query"`

	// Ordering Comma separated list of ordering the results.
	// You may specify reverse orderings by prefixing the field name with '-'.
	//
	// Available fields: username
	Ordering *string `form:"ordering,omitempty" json:"ordering,omitempty"`
}

// ListDomainInvitationsParams defines parameters for ListDomainInvitations.
type ListDomainInvitationsParams struct {
	// Ordering Comma separated list of ordering the results.
	// You may specify reverse orderings by prefixing the field name with '-'.
	//
	// Available fields: created_at,updated_at
	Ordering *string `form:"ordering,omitempty" json:"ordering,omitempty"`
	Offset   *int    `form:"offset,omitempty"   json:"offset,omitempty"`
	Limit    *int    `form:"limit,omitempty"    json:"limit,omitempty"`
}

// JoinDomainByInvitationParams defines parameters for JoinDomainByInvitation.
type JoinDomainByInvitationParams struct {
	InvitationCode string `form:"invitationCode" json:"invitationCode"`
}

// ListProblemSetsParams defines parameters for ListProblemSets.
type ListProblemSetsParams struct {
	// Ordering Comma separated list of ordering the results.
	// You may specify reverse orderings by prefixing the field name with '-'.
	//
	// Available fields: created_at,updated_at
	Ordering *string `form:"ordering,omitempty" json:"ordering,omitempty"`
	Offset   *int    `form:"offset,omitempty"   json:"offset,omitempty"`
	Limit    *int    `form:"limit,omitempty"    json:"limit,omitempty"`
}

// ListProblemsParams defines parameters for ListProblems.
type ListProblemsParams struct {
	// Ordering Comma separated list of ordering the results.
	// You may specify reverse orderings by prefixing the field name with '-'.
	//
	// Available fields: created_at,updated_at
	Ordering *string `form:"ordering,omitempty" json:"ordering,omitempty"`
	Offset   *int    `form:"offset,omitempty"   json:"offset,omitempty"`
	Limit    *int    `form:"limit,omitempty"    json:"limit,omitempty"`
}

// ListProblemConfigCommitsParams defines parameters for ListProblemConfigCommits.
type ListProblemConfigCommitsParams struct {
	// Ordering Comma separated list of ordering the results.
	// You may specify reverse orderings by prefixing the field name with '-'.
	//
	// Available fields: created_at,updated_at
	Ordering *string `form:"ordering,omitempty" json:"ordering,omitempty"`
	Offset   *int    `form:"offset,omitempty"   json:"offset,omitempty"`
	Limit    *int    `form:"limit,omitempty"    json:"limit,omitempty"`
}

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

// ListRecordsInDomainParams defines parameters for ListRecordsInDomain.
type ListRecordsInDomainParams struct {
	// ProblemSet problem set id
	ProblemSet *openapi_types.UUID `form:"problemSet,omitempty" json:"problemSet,omitempty"`

	// Problem problem id
	Problem *openapi_types.UUID `form:"problem,omitempty" json:"problem,omitempty"`

	// SubmitterId submitter uid
	SubmitterId *openapi_types.UUID `form:"submitterId,omitempty" json:"submitterId,omitempty"`

	// Ordering Comma separated list of ordering the results.
	// You may specify reverse orderings by prefixing the field name with '-'.
	//
	// Available fields: created_at,updated_at
	Ordering *string `form:"ordering,omitempty" json:"ordering,omitempty"`
	Offset   *int    `form:"offset,omitempty"   json:"offset,omitempty"`
	Limit    *int    `form:"limit,omitempty"    json:"limit,omitempty"`
}

// ListDomainRolesParams defines parameters for ListDomainRoles.
type ListDomainRolesParams struct {
	// Ordering Comma separated list of ordering the results.
	// You may specify reverse orderings by prefixing the field name with '-'.
	//
	// Available fields: created_at,updated_at
	Ordering *string `form:"ordering,omitempty" json:"ordering,omitempty"`
}

// ListDomainUsersParams defines parameters for ListDomainUsers.
type ListDomainUsersParams struct {
	// Ordering Comma separated list of ordering the results.
	// You may specify reverse orderings by prefixing the field name with '-'.
	//
	// Available fields: created_at,updated_at
	Ordering *string `form:"ordering,omitempty" json:"ordering,omitempty"`
	Offset   *int    `form:"offset,omitempty"   json:"offset,omitempty"`
	Limit    *int    `form:"limit,omitempty"    json:"limit,omitempty"`
}

// ListProblemGroupsParams defines parameters for ListProblemGroups.
type ListProblemGroupsParams struct {
	// Ordering Comma separated list of ordering the results.
	// You may specify reverse orderings by prefixing the field name with '-'.
	//
	// Available fields: created_at,updated_at
	Ordering *string `form:"ordering,omitempty" json:"ordering,omitempty"`
	Offset   *int    `form:"offset,omitempty"   json:"offset,omitempty"`
	Limit    *int    `form:"limit,omitempty"    json:"limit,omitempty"`
}

// AdminCreateJudgerJSONRequestBody defines body for AdminCreateJudger for application/json ContentType.
type AdminCreateJudgerJSONRequestBody = JudgerCreate

// LoginFormdataRequestBody defines body for Login for application/x-www-form-urlencoded ContentType.
type LoginFormdataRequestBody = OAuth2PasswordRequestForm

// RegisterJSONRequestBody defines body for Register for application/json ContentType.
type RegisterJSONRequestBody = UserCreate

// CreateDomainJSONRequestBody defines body for CreateDomain for application/json ContentType.
type CreateDomainJSONRequestBody = DomainCreate

// UpdateDomainJSONRequestBody defines body for UpdateDomain for application/json ContentType.
type UpdateDomainJSONRequestBody = DomainEdit

// CreateDomainInvitationJSONRequestBody defines body for CreateDomainInvitation for application/json ContentType.
type CreateDomainInvitationJSONRequestBody = DomainInvitationCreate

// UpdateDomainInvitationJSONRequestBody defines body for UpdateDomainInvitation for application/json ContentType.
type UpdateDomainInvitationJSONRequestBody = DomainInvitationEdit

// CreateProblemSetJSONRequestBody defines body for CreateProblemSet for application/json ContentType.
type CreateProblemSetJSONRequestBody = ProblemSetCreate

// UpdateProblemSetJSONRequestBody defines body for UpdateProblemSet for application/json ContentType.
type UpdateProblemSetJSONRequestBody = ProblemSetEdit

// AddProblemInProblemSetJSONRequestBody defines body for AddProblemInProblemSet for application/json ContentType.
type AddProblemInProblemSetJSONRequestBody = ProblemSetAddProblem

// UpdateProblemInProblemSetJSONRequestBody defines body for UpdateProblemInProblemSet for application/json ContentType.
type UpdateProblemInProblemSetJSONRequestBody = ProblemSetUpdateProblem

// SubmitSolutionToProblemSetMultipartRequestBody defines body for SubmitSolutionToProblemSet for multipart/form-data ContentType.
type SubmitSolutionToProblemSetMultipartRequestBody = ProblemSolutionSubmit

// CreateProblemJSONRequestBody defines body for CreateProblem for application/json ContentType.
type CreateProblemJSONRequestBody = ProblemCreate

// CloneProblemJSONRequestBody defines body for CloneProblem for application/json ContentType.
type CloneProblemJSONRequestBody = ProblemClone

// UpdateProblemJSONRequestBody defines body for UpdateProblem for application/json ContentType.
type UpdateProblemJSONRequestBody = ProblemEdit

// SubmitSolutionToProblemMultipartRequestBody defines body for SubmitSolutionToProblem for multipart/form-data ContentType.
type SubmitSolutionToProblemMultipartRequestBody = ProblemSolutionSubmit

// UpdateProblemConfigByArchiveMultipartRequestBody defines body for UpdateProblemConfigByArchive for multipart/form-data ContentType.
type UpdateProblemConfigByArchiveMultipartRequestBody = FileUpload

// UpdateProblemConfigJsonJSONRequestBody defines body for UpdateProblemConfigJson for application/json ContentType.
type UpdateProblemConfigJsonJSONRequestBody = ProblemConfigJson

// SubmitCaseByJudgerJSONRequestBody defines body for SubmitCaseByJudger for application/json ContentType.
type SubmitCaseByJudgerJSONRequestBody = RecordCaseSubmit

// SubmitRecordByJudgerJSONRequestBody defines body for SubmitRecordByJudger for application/json ContentType.
type SubmitRecordByJudgerJSONRequestBody = RecordSubmit

// ClaimRecordByJudgerJSONRequestBody defines body for ClaimRecordByJudger for application/json ContentType.
type ClaimRecordByJudgerJSONRequestBody = JudgerClaim

// CreateDomainRoleJSONRequestBody defines body for CreateDomainRole for application/json ContentType.
type CreateDomainRoleJSONRequestBody = DomainRoleCreate

// UpdateDomainRoleJSONRequestBody defines body for UpdateDomainRole for application/json ContentType.
type UpdateDomainRoleJSONRequestBody = DomainRoleEdit

// TransferDomainJSONRequestBody defines body for TransferDomain for application/json ContentType.
type TransferDomainJSONRequestBody = DomainTransfer

// AddDomainUserJSONRequestBody defines body for AddDomainUser for application/json ContentType.
type AddDomainUserJSONRequestBody = DomainUserAdd

// UpdateDomainUserJSONRequestBody defines body for UpdateDomainUser for application/json ContentType.
type UpdateDomainUserJSONRequestBody = DomainUserUpdate

// UpdateCurrentUserJSONRequestBody defines body for UpdateCurrentUser for application/json ContentType.
type UpdateCurrentUserJSONRequestBody = UserEdit

// ChangePasswordJSONRequestBody defines body for ChangePassword for application/json ContentType.
type ChangePasswordJSONRequestBody = UserResetPassword

// AsValidationErrorLoc0 returns the union data inside the ValidationError_Loc_Item as a ValidationErrorLoc0
func (t ValidationError_Loc_Item) AsValidationErrorLoc0() (ValidationErrorLoc0, error) {
	var body ValidationErrorLoc0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromValidationErrorLoc0 overwrites any union data inside the ValidationError_Loc_Item as the provided ValidationErrorLoc0
func (t *ValidationError_Loc_Item) FromValidationErrorLoc0(v ValidationErrorLoc0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeValidationErrorLoc0 performs a merge with any union data inside the ValidationError_Loc_Item, using the provided ValidationErrorLoc0
func (t *ValidationError_Loc_Item) MergeValidationErrorLoc0(v ValidationErrorLoc0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

// AsValidationErrorLoc1 returns the union data inside the ValidationError_Loc_Item as a ValidationErrorLoc1
func (t ValidationError_Loc_Item) AsValidationErrorLoc1() (ValidationErrorLoc1, error) {
	var body ValidationErrorLoc1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromValidationErrorLoc1 overwrites any union data inside the ValidationError_Loc_Item as the provided ValidationErrorLoc1
func (t *ValidationError_Loc_Item) FromValidationErrorLoc1(v ValidationErrorLoc1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeValidationErrorLoc1 performs a merge with any union data inside the ValidationError_Loc_Item, using the provided ValidationErrorLoc1
func (t *ValidationError_Loc_Item) MergeValidationErrorLoc1(v ValidationErrorLoc1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

func (t ValidationError_Loc_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ValidationError_Loc_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}