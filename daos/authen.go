package daos

import (
	"BuildDBGo/dtos"
	"BuildDBGo/models"

	"github.com/jinzhu/gorm"
)

type AuthenDao interface {
	CreateAuthen(record *models.AuthenKey) error
	SearchAuthen(queries dtos.AuthenKey) (models.AuthenKey, error)
}

type authenDaoImpl struct {
	db *gorm.DB
}

func NewAuthenDao(db *gorm.DB) AuthenDao {
	return &authenDaoImpl{db: db}
}

func (dao *authenDaoImpl) CreateAuthen(record *models.AuthenKey) error {
	err := dao.db.Exec("DELETE FROM authenkey;")
	if err != nil {
		//
	}
	return dao.db.Create(record).Error
}

func (dao *authenDaoImpl) SearchAuthen(queries dtos.AuthenKey) (models.AuthenKey, error) {
	result := models.AuthenKey{}
	db := dao.db

	if err := db.Find(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}
