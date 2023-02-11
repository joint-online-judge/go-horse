package schema

// Defines values for ArchiveType.
const (
	Rar     ArchiveType = "rar"
	Tar     ArchiveType = "tar"
	Unknown ArchiveType = "unknown"
	Zip     ArchiveType = "zip"
)

// ArchiveType An enumeration.
type ArchiveType string

// Defines values for ConfigMissing.
const (
	RaiseError ConfigMissing = "raise_error"
	UseDefault ConfigMissing = "use_default"
	UseOld     ConfigMissing = "use_old"
)

// ConfigMissing An enumeration.
type ConfigMissing string

// Defines values for DiffTypeEnum.
const (
	Added    DiffTypeEnum = "added"
	Changed  DiffTypeEnum = "changed"
	Conflict DiffTypeEnum = "conflict"
	Removed  DiffTypeEnum = "removed"
)

// DiffTypeEnum An enumeration.
type DiffTypeEnum string

// Defines values for PathTypeEnum.
const (
	CommonPrefix PathTypeEnum = "common_prefix"
	Object       PathTypeEnum = "object"
)

// PathTypeEnum An enumeration.
type PathTypeEnum string

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

// RecordCaseResult An enumeration.
type RecordCaseResult string

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

// RecordState An enumeration.
type RecordState string

// Defines values for LoginParamsResponseType.
const (
	LoginParamsResponseTypeJson     LoginParamsResponseType = "json"
	LoginParamsResponseTypeRedirect LoginParamsResponseType = "redirect"
)

// LoginParamsResponseType defines parameters for Login.
type LoginParamsResponseType string

// Defines values for LogoutParamsResponseType.
const (
	LogoutParamsResponseTypeJson     LogoutParamsResponseType = "json"
	LogoutParamsResponseTypeRedirect LogoutParamsResponseType = "redirect"
)

// LogoutParamsResponseType defines parameters for Logout.
type LogoutParamsResponseType string

// Defines values for OauthAuthorizeParamsResponseType.
const (
	OauthAuthorizeParamsResponseTypeJson     OauthAuthorizeParamsResponseType = "json"
	OauthAuthorizeParamsResponseTypeRedirect OauthAuthorizeParamsResponseType = "redirect"
)

// OauthAuthorizeParamsResponseType defines parameters for OauthAuthorize.
type OauthAuthorizeParamsResponseType string

// Defines values for RefreshParamsResponseType.
const (
	RefreshParamsResponseTypeJson     RefreshParamsResponseType = "json"
	RefreshParamsResponseTypeRedirect RefreshParamsResponseType = "redirect"
)

// RefreshParamsResponseType defines parameters for Refresh.
type RefreshParamsResponseType string

// Defines values for RegisterParamsResponseType.
const (
	RegisterParamsResponseTypeJson     RegisterParamsResponseType = "json"
	RegisterParamsResponseTypeRedirect RegisterParamsResponseType = "redirect"
)

// RegisterParamsResponseType defines parameters for Register.
type RegisterParamsResponseType string

// Defines values for GetTokenParamsResponseType.
const (
	GetTokenParamsResponseTypeJson     GetTokenParamsResponseType = "json"
	GetTokenParamsResponseTypeRedirect GetTokenParamsResponseType = "redirect"
)

// GetTokenParamsResponseType defines parameters for GetToken.
type GetTokenParamsResponseType string
