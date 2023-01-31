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
	CreatedAt      time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index" json:"updated_at"`
	LoginAt        time.Time `gorm:"column:login_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)"         json:"login_at"`
	ID             uuid.UUID `gorm:"column:id;primaryKey;type:uuid"                                                    json:"id"`
	Username       string    `gorm:"column:username;not null"                                                          json:"username"`
	Email          string    `gorm:"column:email;not null"                                                             json:"email"`
	Gravatar       string    `gorm:"column:gravatar;not null"                                                          json:"gravatar"`
	StudentID      string    `gorm:"column:student_id;not null;index"                                                  json:"student_id"`
	RealName       string    `gorm:"column:real_name;not null;index"                                                   json:"real_name"`
	Role           string    `gorm:"column:role;not null;index"                                                        json:"role"`
	IsActive       bool      `gorm:"column:is_active;not null"                                                         json:"is_active"`
	HashedPassword string    `gorm:"column:hashed_password;not null"                                                   json:"hashed_password"`
	UsernameLower  string    `gorm:"column:username_lower;not null;index:unique"                                       json:"username_lower"`
	EmailLower     string    `gorm:"column:email_lower;not null;index:unique"                                          json:"email_lower"`
	RegisterIP     string    `gorm:"column:register_ip;not null"                                                       json:"register_ip"`
	LoginIP        string    `gorm:"column:login_ip;not null"                                                          json:"login_ip"`
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
