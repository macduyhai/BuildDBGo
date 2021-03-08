package models

type User struct {
	ID           int64  `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY"`
	username     string `gorm:"column:username;type:varchar(128)" json:"username"`
	password     string `gorm:"column:password;type:varchar(128)" json:"password"`
	fullname     string `gorm:"column:fullname;type:varchar(128)" json:"fullname"`
	phone_number string `gorm:"column:phone_number;type:varchar(128)" json:"phone_number"`
	number_video int64  `gorm:"column:number_video;type:int(10)" json:"number_video"`
	total_size   int64  `gorm:"column:total_size;type:int(10)" json:"total_size"`
	max_size     int64  `gorm:"column:max_size;type:int(10)" json:"max_size"`
}

func (User) TableName() string {
	return "users"
}
