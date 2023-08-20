package model

import "time"

type CartCache struct {
	ID       string    `json:"-" gorm:"column:id;primary_key;type:uuid;default:uuid_generate_v4();unique"`
	ItemID   string    `json:"item_id" gorm:"column:item_id"`
	UserID   string    `json:"user_id" gorm:"column:user_id"`
	Quantity int       `json:"quatity" gorm:"column:quantity"`
	Expiry   time.Time `json:"-" gorm:"column:expiry"`
}
