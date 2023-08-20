package model

import "time"

type AccessToken struct {
	ID           string    `json:"-" gorm:"column:id;primary_key;type:uuid;default:uuid_generate_v4();unique"`
	UserID       string    `json:"user_id" gorm:"column:user_id;unique"`
	AccessToken  string    `json:"access_token" gorm:"column:access_token"`
	RefreshToken string    `json:"refresh_token" gorm:"column:refresh_token"`
	Role         string    `json:"role" gorm:"column:role"`
	Active       bool      `json:"active" gorm:"column:active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
