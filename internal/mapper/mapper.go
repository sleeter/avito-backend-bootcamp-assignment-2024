package mapper

import (
	"backend-bootcamp-assignment-2024/internal/model/dto/response"
	"backend-bootcamp-assignment-2024/internal/model/entity"
)

func FlatEntityToFlatResponse(flat *entity.Flat) *response.Flat {
	return &response.Flat{
		Id:      flat.Id,
		HouseId: flat.HouseId,
		Price:   flat.Price,
		Rooms:   flat.Rooms,
		Status:  flat.Status,
	}
}

func HouseEntityToHouseResponse(house *entity.House) *response.House {
	return &response.House{
		Id:        house.Id,
		Address:   house.Address,
		Year:      house.Year,
		Developer: house.Developer,
		CreatedAt: house.CreatedAt,
		UpdateAt:  house.UpdateAt,
	}
}
