package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	CreatedAt      time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index" json:"updatedAt"`
	LoginAt        time.Time `gorm:"column:login_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)"         json:"loginAt"`
	ID             uuid.UUID `gorm:"column:id;primaryKey;type:uuid"                                                    json:"id"`
	Username       string    `gorm:"column:username;not null"                                                          json:"username"`
	Email          string    `gorm:"column:email;not null"                                                             json:"email"`
	Gravatar       string    `gorm:"column:gravatar;not null"                                                          json:"gravatar"`
	StudentID      string    `gorm:"column:student_id;not null;index"                                                  json:"studentId"`
	RealName       string    `gorm:"column:real_name;not null;index"                                                   json:"realName"`
	Role           string    `gorm:"column:role;not null;index"                                                        json:"role"`
	IsActive       bool      `gorm:"column:is_active;not null"                                                         json:"isActive"`
	HashedPassword string    `gorm:"column:hashed_password;not null"                                                   json:"hashedPassword"`
	UsernameLower  string    `gorm:"column:username_lower;not null;index:unique"                                       json:"usernameLower"`
	EmailLower     string    `gorm:"column:email_lower;not null;index:unique"                                          json:"emailLower"`
	RegisterIP     string    `gorm:"column:register_ip;not null"                                                       json:"registerIp"`
	LoginIP        string    `gorm:"column:login_ip;not null"                                                          json:"loginIp"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UsernameLower = strings.ToLower(u.Username)
	u.EmailLower = strings.ToLower(u.Email)
	return nil
}
