package models

import "github.com/google/uuid"

const TableNameUserOauthAccount = "user_oauth_accounts"

// UserOauthAccount mapped from table <user_oauth_accounts>
type UserOauthAccount struct {
	Base
	UserID       uuid.UUID `gorm:"column:user_id;type:uuid"                       json:"userId"`
	User         User      `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	OauthName    string    `gorm:"column:oauth_name;not null"                     json:"oauthName"`
	AccessToken  string    `gorm:"column:access_token;not null"                   json:"accessToken"`
	RefreshToken string    `gorm:"column:refresh_token"                           json:"refreshToken"`
	ExpiresAt    int32     `gorm:"column:expires_at"                              json:"expiresAt"`
	AccountID    string    `gorm:"column:account_id;not null;index"               json:"accountId"`
	AccountName  string    `gorm:"column:account_name;index"                      json:"accountName"`
	AccountEmail string    `gorm:"column:account_email;not null;index"            json:"accountEmail"`
}

// TableName UserOauthAccount's table name
func (*UserOauthAccount) TableName() string {
	return TableNameUserOauthAccount
}
