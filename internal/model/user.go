package model

import "time"

type Role string

const (
	RoleAdmin Role = "MERGE_ROLE_ADMIN"
	RoleUser  Role = "MERGE_ROLE_USER"
)

type User struct {
	UserID    string    `json:"user_id" gorm:"primary_key;type:uuid;default:uuid_generate_v4();unique"`
	Email     string    `json:"email" gorm:"column:email;unique"`
	Password  string    `json:"password" gorm:"column:password"`
	Role      string    `json:"role" gorm:"column:role"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
	Disabled  bool      `json:"-" gorm:"column:disabled;"`
}
