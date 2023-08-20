package model

import "time"

type Item struct {
	ID        string    `json:"id" gorm:"column:id;primary_key;type:uuid;default:uuid_generate_v4();unique"`
	Name      string    `json:"name" gorm:"column:name"`
	Desc      string    `json:"desc" gorm:"column:description"`
	Stock     uint      `json:"stock" gorm:"column:stock"`
	Price     uint      `json:"price" gorm:"column:price"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
	Disabled  bool      `json:"-" gorm:"column:disabled;"`
}
