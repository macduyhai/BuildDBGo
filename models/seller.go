package models

import "time"

type Seller struct {
	ID        int64      `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY"`
	Name      string     `gorm:"column:name"`
	Note      string     `gorm:"column:note"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

func (Seller) TableName() string {
	return "sellers"
}
