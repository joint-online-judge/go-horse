// Package schema provides primitives to interact with the openapi HTTP API.
package schema

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// User defines model for User.
type User struct {
	Gravatar *string            `json:"gravatar,omitempty"`
	Id       openapi_types.UUID `json:"id"`
	IsActive *bool              `json:"isActive,omitempty"`
	Role     *string            `json:"role,omitempty"`
	Username string             `json:"username"`
}

func VerifyPassword(password, hashed_password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	encoded, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(encoded), err
}

// UserCreate defines model for UserCreate.
type UserCreate struct {
	Email          *string `json:"email,omitempty"`
	OauthAccountId *string `json:"oauthAccountId,omitempty"`
	OauthName      *string `json:"oauthName,omitempty"`
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

// UserEdit defines model for UserEdit.
type UserEdit struct {
	Gravatar *string `json:"gravatar,omitempty"`
}

// UserList defines model for UserList.
type UserList struct {
	Count   *int    `json:"count,omitempty"`
	Results *[]User `json:"results,omitempty"`
}

// UserPreview defines model for UserPreview.
type UserPreview struct {
	Gravatar *string            `json:"gravatar,omitempty"`
	Id       openapi_types.UUID `json:"id"`
	Username string             `json:"username"`
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
