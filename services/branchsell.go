package services

import (
	"BuildDBGo/daos"
	"BuildDBGo/dtos"
	"BuildDBGo/models"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type BranchSellService interface {
	AddBranchSell(request dtos.AddbranchRequest) (*dtos.AddbranchResponse, error)
	// AddLabelsToOrder(request dtos.AddLabelRequest) (*dtos.AddorderResponse, error)
	SearchBranch(queries []dtos.SearchBranchSellQuery) ([]dtos.BranchSell, error)
	// AddShippingTime(request dtos.AddShippingTimeRequest) (*dtos.AddorderResponse, error)
	// MakeCompleted(orderNumber string) (*dtos.AddorderResponse, error)
	Detete(branchname string) error
	Updates(request dtos.BranchSell) (*dtos.AddbranchResponse, error)
}

type branchSellServiceImpl struct {
	dao daos.BranchSellDao
}

func NewBranchSellService(dao daos.BranchSellDao) BranchSellService {
	return &branchSellServiceImpl{
		dao: dao,
	}
}

func (service *branchSellServiceImpl) mapperDtossToModelBranchSell(input dtos.BranchSell) models.BranchSell {
	return models.BranchSell{
		Name: input.Name,
		Note: input.Note,
	}
}
func (service *branchSellServiceImpl) mapperModelToDtossBranchSell(input models.BranchSell) dtos.BranchSell {
	return dtos.BranchSell{
		Name: input.Name,
		Note: input.Note,
	}
}

func (service *branchSellServiceImpl) AddBranchSell(request dtos.AddbranchRequest) (*dtos.AddbranchResponse, error) {
	recordSuccess := make([]string, 0)
	recordFail := make([]string, 0)
	for _, branchsell := range request.BranchSells {
		branchsell.Name = strings.ToLower(branchsell.Name)
		branchsell.Name = strings.Replace(branchsell.Name, " ", "", -1)
		record := service.mapperDtossToModelBranchSell(branchsell)
		err := service.dao.Create(&record)
		if err != nil {
			recordFail = append(recordFail, branchsell.Name)
		} else {
			recordSuccess = append(recordSuccess, branchsell.Name)
		}
	}

	result := dtos.AddbranchResponse{
		RecordsFailes:  recordFail,
		RecordsSuccess: recordSuccess,
	}

	return &result, nil
}

// func (service *branchSellServiceImpl) SearchBranchforName(name string) bool {
// 	query := []dtos.SearchBranchSellQuery{}
// 	query[0].Name = name
// 	record, _ := service.dao.SearchBranch(query)

// 	if record[0].Name == name {
// 		return true
// 	} else {
// 		return false
// 	}

// }

func (service *branchSellServiceImpl) SearchBranch(queries []dtos.SearchBranchSellQuery) ([]dtos.BranchSell, error) {
	records, _ := service.dao.SearchBranch(queries)
	result := make([]dtos.BranchSell, 0)
	name := ""

	for _, query := range queries {
		if query.Key == "name=?" {
			name = fmt.Sprintf("%v", query.Value)
		}
	}
	for _, record := range records {
		if name != "" {
			if record.Name == name {
				result = append(result, service.mapperModelToDtossBranchSell(record))
			}
		} else {
			result = append(result, service.mapperModelToDtossBranchSell(record))
		}
	}

	return result, nil
}

func (service *branchSellServiceImpl) Detete(branchname string) error {
	fmt.Println(branchname)
	return service.dao.Delete(branchname)
}
func (service *branchSellServiceImpl) Updates(request dtos.BranchSell) (*dtos.AddbranchResponse, error) {
	record := service.mapperDtossToModelBranchSell(request)
	err := service.dao.Updates(&record)
	if err != nil {
		return nil, errors.Wrap(err, "update record error")
	}

	return &dtos.AddbranchResponse{ID: record.ID}, nil
}
