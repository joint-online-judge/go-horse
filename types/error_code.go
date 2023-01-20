package types

import "fmt"

// ErrorCode An enumeration.
type ErrorCode string

// Defines values for ErrorCode.
const (
	APINotImplementedError             ErrorCode = "APINotImplementedError"
	DeleteProblemBadRequestError       ErrorCode = "DeleteProblemBadRequestError"
	DomainInvitationBadRequestError    ErrorCode = "DomainInvitationBadRequestError"
	DomainInvitationExpiredError       ErrorCode = "DomainInvitationExpiredError"
	DomainNotFoundError                ErrorCode = "DomainNotFoundError"
	DomainNotOwnerError                ErrorCode = "DomainNotOwnerError"
	DomainNotRootError                 ErrorCode = "DomainNotRootError"
	DomainRoleNotFoundError            ErrorCode = "DomainRoleNotFoundError"
	DomainRoleNotUniqueError           ErrorCode = "DomainRoleNotUniqueError"
	DomainRoleReadOnlyError            ErrorCode = "DomainRoleReadOnlyError"
	DomainRoleUsedError                ErrorCode = "DomainRoleUsedError"
	DomainUserNotFoundError            ErrorCode = "DomainUserNotFoundError"
	Error                              ErrorCode = "Error"
	FileDownloadError                  ErrorCode = "FileDownloadError"
	FileSystemError                    ErrorCode = "FileSystemError"
	FileUpdateError                    ErrorCode = "FileUpdateError"
	FileValidationError                ErrorCode = "FileValidationError"
	IllegalFieldError                  ErrorCode = "IllegalFieldError"
	IntegrityError                     ErrorCode = "IntegrityError"
	InternalServerError                ErrorCode = "InternalServerError"
	InvalidUrlError                    ErrorCode = "InvalidUrlError"
	LockError                          ErrorCode = "LockError"
	ProblemConfigJsonNotFoundError     ErrorCode = "ProblemConfigJsonNotFoundError"
	ProblemConfigNotFoundError         ErrorCode = "ProblemConfigNotFoundError"
	ProblemGroupNotFoundError          ErrorCode = "ProblemGroupNotFoundError"
	ProblemNotFoundError               ErrorCode = "ProblemNotFoundError"
	ProblemSetNotFoundError            ErrorCode = "ProblemSetNotFoundError"
	RecordNotFoundError                ErrorCode = "RecordNotFoundError"
	ScoreboardHiddenBadRequestError    ErrorCode = "ScoreboardHiddenBadRequestError"
	Success                            ErrorCode = "Success"
	UnknownFieldError                  ErrorCode = "UnknownFieldError"
	UnsupportedLanguageError           ErrorCode = "UnsupportedLanguageError"
	UserAlreadyInDomainBadRequestError ErrorCode = "UserAlreadyInDomainBadRequestError"
	UserNotFoundError                  ErrorCode = "UserNotFoundError"
	UserNotJudgerError                 ErrorCode = "UserNotJudgerError"
	UserRegisterError                  ErrorCode = "UserRegisterError"
	UsernamePasswordError              ErrorCode = "UsernamePasswordError"
)

type BizError struct {
	ErrorCode ErrorCode `json:"errorCode"`
	ErrorMsg  *string   `json:"errorMsg,omitempty"`
}

func (e BizError) Error() string {
	return fmt.Sprintf("BizError: %s, %s", e.ErrorCode, *e.ErrorMsg)
}

