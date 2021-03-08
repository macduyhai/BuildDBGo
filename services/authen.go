package services

import (
	"order-ops/daos"
	"order-ops/dtos"
	"order-ops/models"
)

type AuthenService interface {
	AddAuthen(request dtos.AuthenKey) (*dtos.AuthenKey, error)
	SearchAuthen() (dtos.AuthenKey, error)
}

type authenServiceImpl struct {
	dao daos.AuthenDao
}

func NewAuthenService(dao daos.AuthenDao) AuthenService {
	return &authenServiceImpl{
		dao: dao,
	}
}

func (service *authenServiceImpl) AddAuthen(request dtos.AuthenKey) (*dtos.AuthenKey, error) {
	recordSuccess := make([]string, 0)
	recordFail := make([]string, 0)

	query := models.AuthenKey{}
	query.Key = request.Key
	query.CreatedAt = request.CreatedAt
	err := service.dao.CreateAuthen(&query)
	if err != nil {
		recordFail = append(recordFail, query.Key)
	} else {
		recordSuccess = append(recordSuccess, query.Key)
	}

	result := dtos.AuthenKey{
		Key:       query.Key,
		CreatedAt: query.CreatedAt,
	}

	return &result, nil
}

func (service *authenServiceImpl) SearchAuthen() (dtos.AuthenKey, error) {
	queries := dtos.AuthenKey{}
	record, _ := service.dao.SearchAuthen(queries)
	result := dtos.AuthenKey{}
	if record.Key != "" {
		result.Key = record.Key
		result.CreatedAt = record.CreatedAt
	}

	return result, nil
}
