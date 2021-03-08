package models

import "time"

type TypeProduct struct {
	ID        int64      `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY"`
	Name      string     `gorm:"column:name"`
	Width     int32      `gorm:"column:width"`
	Height    int32      `gorm:"column:height"`
	Weight    int32      `gorm:"column:weight"`
	Length    int32      `gorm:"column:length"`
	Note      string     `gorm:"column:note"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

func (TypeProduct) TableName() string {
	return "typeproducts"
}
