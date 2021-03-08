package daos

import (
	"BuildDBGo/dtos"
	"BuildDBGo/models"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type BranchSellDao interface {
	Create(record *models.BranchSell) error
	Updates(record *models.BranchSell) error
	SearchBranch(queries []dtos.SearchBranchSellQuery) ([]models.BranchSell, error)
	// GetByOrderNumber(orderNumber string) (*models.Order, error)
	Delete(branchname string) error
}

type branchSellDaoImpl struct {
	db *gorm.DB
}

func NewBranchSellDao(db *gorm.DB) BranchSellDao {
	return &branchSellDaoImpl{db: db}
}

func (dao *branchSellDaoImpl) Create(record *models.BranchSell) error {
	existedRecord, err := dao.GetByBranchName(record.Name)
	if err != nil {
		return dao.db.Create(record).Error
	}

	record.ID = existedRecord.ID
	return dao.db.Model(&existedRecord).Where("id=?", existedRecord.ID).Updates(record).Error

}
func (dao *branchSellDaoImpl) GetByBranchName(branchName string) (*models.BranchSell, error) {
	var result models.BranchSell
	if err := dao.db.Where("name=?", branchName).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
func (dao *branchSellDaoImpl) Updates(record *models.BranchSell) error {
	existedRecord, err := dao.GetByBranchName(record.Name)
	if err != nil {
		return errors.Wrap(err, "get existed record error")
	}

	record.ID = existedRecord.ID
	return dao.db.Model(&existedRecord).Where("id=?", existedRecord.ID).Updates(record).Error
}

func (dao *branchSellDaoImpl) SearchBranch(queries []dtos.SearchBranchSellQuery) ([]models.BranchSell, error) {
	result := make([]models.BranchSell, 0)
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
func (dao *branchSellDaoImpl) Delete(branchname string) error {
	var result models.BranchSell
	if err := dao.db.Where("name=?", branchname).Delete(&result).Error; err != nil {
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
