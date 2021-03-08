package daos

import (
	"BuildDBGo/dtos"
	"BuildDBGo/models"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type SellerDao interface {
	Create(record *models.Seller) error
	Updates(record *models.Seller) error
	SearchSeller(queries []dtos.SearchSellerQuery) ([]models.Seller, error)
	// GetByOrderNumber(orderNumber string) (*models.Order, error)
	Delete(Sellername string) error
}

type sellerDaoImpl struct {
	db *gorm.DB
}

func NewSellerDao(db *gorm.DB) SellerDao {
	return &sellerDaoImpl{db: db}
}

func (dao *sellerDaoImpl) Create(record *models.Seller) error {
	existedRecord, err := dao.GetBySellerName(record.Name)
	if err != nil {
		return dao.db.Create(record).Error
	}

	record.ID = existedRecord.ID
	return dao.db.Model(&existedRecord).Where("id=?", existedRecord.ID).Updates(record).Error

}

func (dao *sellerDaoImpl) GetBySellerName(sellerName string) (*models.Seller, error) {
	var result models.Seller
	if err := dao.db.Where("name=?", sellerName).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
func (dao *sellerDaoImpl) Updates(record *models.Seller) error {
	existedRecord, err := dao.GetBySellerName(record.Name)
	if err != nil {
		return errors.Wrap(err, "get existed record error")
	}

	record.ID = existedRecord.ID
	return dao.db.Model(&existedRecord).Where("id=?", existedRecord.ID).Updates(record).Error
}

func (dao *sellerDaoImpl) SearchSeller(queries []dtos.SearchSellerQuery) ([]models.Seller, error) {
	result := make([]models.Seller, 0)
	db := dao.db
	for _, query := range queries {
		if query.Key == "name=?" {
			continue
		}

		if query.Value != nil {
			db = db.Where(query.Key, query.Value)
		} else {
			db = db.Where(query.Key)
		}
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, nil
	}

	return result, nil
}

func (dao *sellerDaoImpl) Delete(Sellername string) error {
	var result models.Seller
	if err := dao.db.Where("name=?", Sellername).Delete(&result).Error; err != nil {
		return err
	}

	return nil

	// return dao.db.Where("name=?", branchname).Delete(&result).Error
}

// func (dao *orderDaoImpl) GetByOrderNumber(orderNumber string) (*models.Order, error) {
// 	var result models.Order
// 	if err := dao.db.Where("order_number=?", orderNumber).First(&result).Error; err != nil {
// 		return nil, err
// 	}

// 	return &result, nil
// }

// func (dao *orderDaoImpl) Delete(orderNumber string) error {
// 	var result models.Order
// 	return dao.db.Where("order_number=?", orderNumber).Delete(&result).Error
// }
