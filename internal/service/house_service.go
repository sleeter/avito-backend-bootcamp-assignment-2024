//go:generate mockgen -source ./house_service.go -destination=./mocks/house_service.go -package=mock_service

package service

import (
	"context"
	"time"

	"backend-bootcamp-assignment-2024/internal/mapper"
	"backend-bootcamp-assignment-2024/internal/model/dto/request"
	"backend-bootcamp-assignment-2024/internal/model/dto/response"
	"backend-bootcamp-assignment-2024/internal/model/entity"
)

type HouseRepository interface {
	CreateHouse(ctx context.Context, house request.House) (*entity.House, error)
	UpdateHouse(ctx context.Context, houseId int32, updateTime time.Time) error
}

type HouseService struct {
	Repository         HouseRepository
	TransactionManager TransactionManager
}

func NewHouseService(r HouseRepository, manager TransactionManager) HouseService {
	return HouseService{
		Repository:         r,
		TransactionManager: manager,
	}
}

func (s *HouseService) CreateHouse(ctx context.Context, house request.House) (*response.House, error) {
	res, err := s.Repository.CreateHouse(ctx, house)
	if err != nil {
		return nil, err
	}
	return mapper.HouseEntityToHouseResponse(res), nil
}

func (s *HouseService) UpdateHouse(ctx context.Context, houseId int32, updateTime time.Time) error {
	return s.Repository.UpdateHouse(ctx, houseId, updateTime)
}
