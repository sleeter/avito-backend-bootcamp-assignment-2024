package service

import (
	"backend-bootcamp-assignment-2024/internal/model/dto/request"
	"backend-bootcamp-assignment-2024/internal/model/entity"
	"context"
	"time"
)

type HouseRepository interface {
	CreateHouse(ctx context.Context, house request.House) (entity.House, error)
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
