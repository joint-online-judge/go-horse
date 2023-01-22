package models

import (
	"time"
)

const TableNameUserOauthAccount = "user_oauth_accounts"

// UserOauthAccount mapped from table <user_oauth_accounts>
type UserOauthAccount struct {
	CreatedAt    time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"updated_at"`
	UserID       string    `gorm:"column:user_id"                                                              json:"user_id"`
	ID           string    `gorm:"column:id;primaryKey"                                                        json:"id"`
	OauthName    string    `gorm:"column:oauth_name;not null"                                                  json:"oauth_name"`
	AccessToken  string    `gorm:"column:access_token;not null"                                                json:"access_token"`
	RefreshToken string    `gorm:"column:refresh_token"                                                        json:"refresh_token"`
	ExpiresAt    int32     `gorm:"column:expires_at"                                                           json:"expires_at"`
	AccountID    string    `gorm:"column:account_id;not null"                                                  json:"account_id"`
	AccountName  string    `gorm:"column:account_name"                                                         json:"account_name"`
	AccountEmail string    `gorm:"column:account_email;not null"                                               json:"account_email"`
}

// TableName UserOauthAccount's table name
func (*UserOauthAccount) TableName() string {
	return TableNameUserOauthAccount
}
