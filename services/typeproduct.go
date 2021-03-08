package services

import (
	"BuildDBGo/daos"
	"BuildDBGo/dtos"
	"BuildDBGo/models"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type TypeProductService interface {
	AddTypeProduct(request dtos.AddtypeRequest) (*dtos.AddtypeResponse, error)
	// AddLabelsToOrder(request dtos.AddLabelRequest) (*dtos.AddorderResponse, error)
	SearchType(queries []dtos.SearchTypeProductQuery) ([]dtos.TypeProduct, error)
	// AddShippingTime(request dtos.AddShippingTimeRequest) (*dtos.AddorderResponse, error)
	// MakeCompleted(orderNumber string) (*dtos.AddorderResponse, error)
	Detete(Typename string) error
	Updates(request dtos.TypeProduct) (*dtos.AddtypeResponse, error)
}

type typeProductServiceImpl struct {
	dao daos.TypeProductDao
}

func NewTypeProductService(dao daos.TypeProductDao) TypeProductService {
	return &typeProductServiceImpl{
		dao: dao,
	}
}

func (service *typeProductServiceImpl) mapperDtossToModelTypeProduct(input dtos.TypeProduct) models.TypeProduct {
	return models.TypeProduct{
		Name:   input.Name,
		Width:  input.Width,
		Height: input.Height,
		Weight: input.Weight,
		Length: input.Length,
		Note:   input.Note,
	}
}
func (service *typeProductServiceImpl) mapperModelToDtossTypeProduct(input models.TypeProduct) dtos.TypeProduct {
	return dtos.TypeProduct{
		Name:   input.Name,
		Width:  input.Width,
		Height: input.Height,
		Weight: input.Weight,
		Length: input.Length,
		Note:   input.Note,
	}
}

// func (service *orderServiceImpl) mapperDtossToModelOrderAddLable(input dtos.AddLabelRequest) models.Order {
// 	return models.Order{
// 		OrderNumber:    input.OrderNumber,
// 		TrackingNumber: input.LableDetails.TrackingNumber,
// 		URL:            input.LableDetails.URL,
// 		PartnerTrackingNumber: input.LableDetails.PartnerTrackingNumber,
// 	}
// }

func (service *typeProductServiceImpl) AddTypeProduct(request dtos.AddtypeRequest) (*dtos.AddtypeResponse, error) {
	recordSuccess := make([]string, 0)
	recordFail := make([]string, 0)
	for _, typeproduct := range request.TypeProducts {
		typeproduct.Name = strings.ToLower(typeproduct.Name)
		typeproduct.Name = strings.Replace(typeproduct.Name, " ", "", -1)
		record := service.mapperDtossToModelTypeProduct(typeproduct)
		err := service.dao.Create(&record)
		if err != nil {
			recordFail = append(recordFail, typeproduct.Name)
		} else {
			recordSuccess = append(recordSuccess, typeproduct.Name)
		}
	}

	result := dtos.AddtypeResponse{
		RecordsFailes:  recordFail,
		RecordsSuccess: recordSuccess,
	}

	return &result, nil
}

func (service *typeProductServiceImpl) SearchType(queries []dtos.SearchTypeProductQuery) ([]dtos.TypeProduct, error) {
	records, _ := service.dao.SearchType(queries)
	result := make([]dtos.TypeProduct, 0)
	name := ""

	for _, query := range queries {
		if query.Key == "name=?" {
			name = fmt.Sprintf("%v", query.Value)
		}
	}
	for _, record := range records {
		if name != "" {
			if record.Name == name {
				result = append(result, service.mapperModelToDtossTypeProduct(record))
			}
		} else {
			result = append(result, service.mapperModelToDtossTypeProduct(record))
		}
	}

	return result, nil
}
func (service *typeProductServiceImpl) Detete(Typename string) error {
	return service.dao.Delete(Typename)
}
func (service *typeProductServiceImpl) Updates(request dtos.TypeProduct) (*dtos.AddtypeResponse, error) {
	record := service.mapperDtossToModelTypeProduct(request)
	err := service.dao.Updates(&record)
	if err != nil {
		return nil, errors.Wrap(err, "update record error")
	}

	return &dtos.AddtypeResponse{ID: record.ID}, nil
}
