package schemas

type StandardResp[T any] struct {
	BizError
	Data T `json:"data,omitempty"`
}

func NewEmptyResp(errorCode ErrorCode, errorMsg ...string) StandardResp[any] {
	return StandardResp[any]{
		BizError: NewBizError(errorCode, errorMsg...),
		Data:     nil,
	}
}

type ListResp[T any] struct {
	Count   int64 `json:"count"`
	Results []T   `json:"results" validate:"dive"`
}

type StandardListResp[T any] StandardResp[ListResp[T]]

func NewListResp[T any](count int64, results ...[]T) ListResp[T] {
	if count == 0 {
		return ListResp[T]{Count: count, Results: []T{}}
	}
	return ListResp[T]{Count: count, Results: results[0]}
}

type NonStandardResp struct {
	Data any
}

func NewNonStandardResp(data any) NonStandardResp {
	return NonStandardResp{Data: data}
}

// AuthTokensResp defines model for AuthTokensResp.
type AuthTokensResp struct {
	Data *AuthTokens `json:"data,omitempty"`
	BizError
}

// AuthTokensWithLakefsResp defines model for AuthTokensWithLakefsResp.
type AuthTokensWithLakefsResp struct {
	Data *AuthTokensWithLakefs `json:"data,omitempty"`
	BizError
}

// DiffListResp defines model for DiffListResp.
type DiffListResp struct {
	Data *DiffList `json:"data,omitempty"`
	BizError
}

// DomainDetailResp defines model for DomainDetailResp.
type DomainDetailResp struct {
	Data *DomainDetail `json:"data,omitempty"`
	BizError
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

// DomainListResp defines model for DomainListResp.
type DomainListResp struct {
	Data *DomainList `json:"data,omitempty"`
	BizError
}

// DomainResp defines model for DomainResp.
type DomainResp struct {
	Data *Domain `json:"data,omitempty"`
	BizError
}

// DomainRoleDetailResp defines model for DomainRoleDetailResp.
type DomainRoleDetailResp struct {
	Data *DomainRoleDetail `json:"data,omitempty"`
	BizError
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

// DomainTagListResp defines model for DomainTagListResp.
type DomainTagListResp struct {
	Data *DomainTagList `json:"data,omitempty"`
	BizError
}

// DomainUserPermissionResp defines model for DomainUserPermissionResp.
type DomainUserPermissionResp struct {
	Data *DomainUserPermission `json:"data,omitempty"`
	BizError
}

// JWTAccessTokenResp defines model for JWTAccessTokenResp.
type JWTAccessTokenResp struct {
	Data *JWTAccessToken `json:"data,omitempty"`
	BizError
}

// JudgerCredentialsResp defines model for JudgerCredentialsResp.
type JudgerCredentialsResp struct {
	Data *JudgerCredentials `json:"data,omitempty"`
	BizError
}

// JudgerDetailListResp defines model for JudgerDetailListResp.
type JudgerDetailListResp struct {
	Data *JudgerDetailList `json:"data,omitempty"`
	BizError
}

// OAuth2ClientListResp defines model for OAuth2ClientListResp.
type OAuth2ClientListResp struct {
	Data *OAuth2ClientList `json:"data,omitempty"`
	BizError
}

// ObjectStatsListResp defines model for ObjectStatsListResp.
type ObjectStatsListResp struct {
	Data *ObjectStatsList `json:"data,omitempty"`
	BizError
}

// ProblemConfigDataDetailResp defines model for ProblemConfigDataDetailResp.
type ProblemConfigDataDetailResp struct {
	Data *ProblemConfigDataDetail `json:"data,omitempty"`
	BizError
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

// ProblemDetailResp defines model for ProblemDetailResp.
type ProblemDetailResp struct {
	Data *ProblemDetail `json:"data,omitempty"`
	BizError
}

// ProblemDetailWithLatestRecordResp defines model for ProblemDetailWithLatestRecordResp.
type ProblemDetailWithLatestRecordResp struct {
	Data *ProblemDetailWithLatestRecord `json:"data,omitempty"`
	BizError
}

// ProblemGroupListResp defines model for ProblemGroupListResp.
type ProblemGroupListResp struct {
	Data *ProblemGroupList `json:"data,omitempty"`
	BizError
}

// ProblemListResp defines model for ProblemListResp.
type ProblemListResp struct {
	Data *ProblemList `json:"data,omitempty"`
	BizError
}

// ProblemResp defines model for ProblemResp.
type ProblemResp struct {
	Data *Problem `json:"data,omitempty"`
	BizError
}

// ProblemSetDetailResp defines model for ProblemSetDetailResp.
type ProblemSetDetailResp struct {
	Data *ProblemSetDetail `json:"data,omitempty"`
	BizError
}

// ProblemSetListResp defines model for ProblemSetListResp.
type ProblemSetListResp struct {
	Data *ProblemSetList `json:"data,omitempty"`
	BizError
}

// ProblemSetResp defines model for ProblemSetResp.
type ProblemSetResp struct {
	Data *ProblemSet `json:"data,omitempty"`
	BizError
}

// ProblemWithLatestRecordListResp defines model for ProblemWithLatestRecordListResp.
type ProblemWithLatestRecordListResp struct {
	Data *ProblemWithLatestRecordList `json:"data,omitempty"`
	BizError
}

// RecordDetailResp defines model for RecordDetailResp.
type RecordDetailResp struct {
	Data *RecordDetail `json:"data,omitempty"`
	BizError
}

// RecordListDetailListResp defines model for RecordListDetailListResp.
type RecordListDetailListResp struct {
	Data *RecordListDetailList `json:"data,omitempty"`
	BizError
}

// RecordResp defines model for RecordResp.
type RecordResp struct {
	Data *Record `json:"data,omitempty"`
	BizError
}

// RedirectResp defines model for RedirectResp.
type RedirectResp struct {
	Data *Redirect `json:"data,omitempty"`
	BizError
}

// UserDetailResp defines model for UserDetailResp.
type UserDetailResp struct {
	Data *UserDetail `json:"data,omitempty"`
	BizError
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

// UserListResp defines model for UserListResp.
type UserListResp struct {
	Data *UserList `json:"data,omitempty"`
	BizError
}

// UserPreviewResp defines model for UserPreviewResp.
type UserPreviewResp struct {
	Data *UserPreview `json:"data,omitempty"`
	BizError
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
	Msg  string `json:"msg"`
	Type string `json:"type"`
}

// HTTPValidationError defines model for HTTPValidationError.
type HTTPValidationError struct {
	Detail *[]ValidationError `json:"detail,omitempty"`
}

type ForbiddenResp struct {
	Detail string `json:"detail"`
}
