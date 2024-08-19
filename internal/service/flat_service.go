package service

import (
	"backend-bootcamp-assignment-2024/internal/mapper"
	"backend-bootcamp-assignment-2024/internal/model/dto/request"
	"backend-bootcamp-assignment-2024/internal/model/dto/response"
	"backend-bootcamp-assignment-2024/internal/model/entity"
	"context"
	"time"
)

type FlatRepository interface {
	GetFlatsByHouseId(ctx context.Context, houseId int32, isModerator bool) ([]entity.Flat, error)
	CreateFlat(ctx context.Context, flat request.CreateFlat) (*entity.Flat, error)
	UpdateFlatStatus(ctx context.Context, flat request.UpdateFlat) (*entity.Flat, error)
}

type FlatService struct {
	HouseRepository    HouseRepository
	FlatRepository     FlatRepository
	TransactionManager TransactionManager
}

func NewFlatService(fr FlatRepository, hr HouseRepository, manager TransactionManager) FlatService {
	return FlatService{
		FlatRepository:     fr,
		HouseRepository:    hr,
		TransactionManager: manager,
	}
}

func (s *FlatService) AddFlat(ctx context.Context, flat request.CreateFlat) (*response.Flat, error) {
	flat.Status = entity.FLATSTATUS_CREATED
	if flat.Rooms == 0 {
		flat.Rooms = 1
	}

	var res entity.Flat
	err := s.TransactionManager.Tx(ctx, func(ctx context.Context) error {
		var err error
		res1, err := s.FlatRepository.CreateFlat(ctx, flat)

		res = *res1

		if err != nil {
			return err
		}

		err = s.HouseRepository.UpdateHouse(ctx, flat.HouseId, time.Now())
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return mapper.FlatEntityToFlatResponse(&res), nil
}

func (s *FlatService) UpdateFlat(ctx context.Context, flat request.UpdateFlat) (*response.Flat, error) {
	res, err := s.FlatRepository.UpdateFlatStatus(ctx, flat)
	if err != nil {
		return nil, err
	}
	return mapper.FlatEntityToFlatResponse(res), nil
}

func (s *FlatService) GetFlats(ctx context.Context, houseId int32, isModerator bool) ([]response.Flat, error) {
	var flats []response.Flat
	res, err := s.FlatRepository.GetFlatsByHouseId(ctx, houseId, isModerator)
	if err != nil {
		return nil, err
	}
	for _, flat := range res {
		flatResponse := mapper.FlatEntityToFlatResponse(&flat)
		flats = append(flats, *flatResponse)
	}
	return flats, nil
}
