package daos

import (
	"BuildDBGo/models"

	"github.com/jinzhu/gorm"
)

type UserDao interface {
	Create(record *models.User) error
	// Updates(record *models.Order) error
	// Create_Item(record *models.Item) error
	// Search(queries []dtos.SearchQuery) ([]models.Order, error)
	// SearchItems(queries []dtos.SearchItemsQuery) ([]models.Item, error)
	// GetByOrderNumber(orderNumber string) (*models.Order, error)
	// Delete(orderNumber string) error
}

type userDaoImpl struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &userDaoImpl{db: db}
}

func (dao *userDaoImpl) Create(record *models.User) error {
	// fmt.Println(record)
	return dao.db.Create(record).Error
}

// func (dao *orderDaoImpl) Create_Item(record *models.Item) error {
// 	// fmt.Println(record)
// 	//log.Println(record)
// 	return dao.db.Create(record).Error
// }
// func (dao *orderDaoImpl) Updates(record *models.Order) error {
// 	existedRecord, err := dao.GetByOrderNumber(record.OrderNumber)
// 	if err != nil {
// 		return errors.Wrap(err, "get existed record error")
// 	}

// 	record.ID = existedRecord.ID
// 	return dao.db.Model(&existedRecord).Where("id=?", existedRecord.ID).Updates(record).Error
// }

// //
// // func (dao *orderDaoImpl) OrderforDay(timeStart time.Time) (val int64) {

// // }
// func (dao *orderDaoImpl) SearchItems(queries []dtos.SearchItemsQuery) ([]models.Item, error) {
// 	result := make([]models.Item, 0)
// 	db := dao.db
// 	for _, query := range queries {

// 		if query.Value != nil {
// 			db = db.Where(query.Key, query.Value)
// 		} else {
// 			db = db.Where(query.Key)
// 		}
// 	}

// 	if err := db.Find(&result).Error; err != nil {
// 		return nil, nil
// 	}

// 	return result, nil
// }

// func (dao *orderDaoImpl) Search(queries []dtos.SearchQuery) ([]models.Order, error) {
// 	result := make([]models.Order, 0)
// 	db := dao.db
// 	for _, query := range queries {
// 		if query.Key == "status=?" {
// 			continue
// 		}

// 		if query.Value != nil {
// 			db = db.Where(query.Key, query.Value)
// 		} else {
// 			db = db.Where(query.Key)
// 		}
// 	}

// 	if err := db.Find(&result).Error; err != nil {
// 		return nil, nil
// 	}

// 	return result, nil
// }

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
