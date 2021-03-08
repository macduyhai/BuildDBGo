package models

import "time"

type AuthenKey struct {
	ID        int64      `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY"`
	Key      string     `gorm:"column:key"`
	CreatedAt *time.Time `gorm:"column:created_at"`
}

func (AuthenKey) TableName() string {
	return "authenkey"
}
