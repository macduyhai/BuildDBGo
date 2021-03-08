package services

import (
	"BuildDBGo/daos"
	"BuildDBGo/dtos"
	"BuildDBGo/models"
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
)

type SellerService interface {
	AddSeller(request dtos.AddsellerRequest) (*dtos.AddsellerResponse, error)
	// AddLabelsToOrder(request dtos.AddLabelRequest) (*dtos.AddorderResponse, error)
	SearchSeller(queries []dtos.SearchSellerQuery) ([]dtos.Seller, error)
	// AddShippingTime(request dtos.AddShippingTimeRequest) (*dtos.AddorderResponse, error)
	// MakeCompleted(orderNumber string) (*dtos.AddorderResponse, error)
	Detete(Sellername string) error
	Updates(request dtos.Seller) (*dtos.AddsellerResponse, error)
}

type sellerServiceImpl struct {
	dao daos.SellerDao
}

func NewSellerService(dao daos.SellerDao) SellerService {
	return &sellerServiceImpl{
		dao: dao,
	}
}

func (service *sellerServiceImpl) mapperDtossToModelSeller(input dtos.Seller) models.Seller {
	return models.Seller{
		Name: input.Name,
		Note: input.Note,
	}
}

func (service *sellerServiceImpl) mapperModelToDtossSeller(input models.Seller) dtos.Seller {
	return dtos.Seller{
		Name: input.Name,
		Note: input.Note,
	}
}

func (service *sellerServiceImpl) AddSeller(request dtos.AddsellerRequest) (*dtos.AddsellerResponse, error) {
	recordSuccess := make([]string, 0)
	recordFail := make([]string, 0)
	for _, seller := range request.Sellers {
		seller.Name = strings.ToLower(seller.Name)
		seller.Name = strings.Replace(seller.Name, " ", "", -1)
		log.Println(seller.Name)
		record := service.mapperDtossToModelSeller(seller)
		err := service.dao.Create(&record)
		if err != nil {
			recordFail = append(recordFail, seller.Name)
		} else {
			recordSuccess = append(recordSuccess, seller.Name)
		}
	}

	result := dtos.AddsellerResponse{
		RecordsFailes:  recordFail,
		RecordsSuccess: recordSuccess,
	}

	return &result, nil
}

func (service *sellerServiceImpl) SearchSeller(queries []dtos.SearchSellerQuery) ([]dtos.Seller, error) {
	records, _ := service.dao.SearchSeller(queries)
	result := make([]dtos.Seller, 0)
	name := ""

	for _, query := range queries {
		if query.Key == "name=?" {
			name = fmt.Sprintf("%v", query.Value)
		}
	}
	for _, record := range records {
		if name != "" {
			if record.Name == name {
				result = append(result, service.mapperModelToDtossSeller(record))
			}
		} else {
			result = append(result, service.mapperModelToDtossSeller(record))
		}
	}

	return result, nil
}

func (service *sellerServiceImpl) Detete(Sellername string) error {
	return service.dao.Delete(Sellername)
}

//
func (service *sellerServiceImpl) Updates(request dtos.Seller) (*dtos.AddsellerResponse, error) {
	record := service.mapperDtossToModelSeller(request)
	err := service.dao.Updates(&record)
	if err != nil {
		return nil, errors.Wrap(err, "update record error")
	}

	return &dtos.AddsellerResponse{ID: record.ID}, nil
}
