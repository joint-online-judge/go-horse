// Package schema provides primitives to interact with the openapi HTTP API.
package schema

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

// SubmitSolutionToProblemInProblemSetMultipartRequestBody defines body for SubmitSolutionToProblemInProblemSet for multipart/form-data ContentType.
type SubmitSolutionToProblemInProblemSetMultipartRequestBody = ProblemSolutionSubmit

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
