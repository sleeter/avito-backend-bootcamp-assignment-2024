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
