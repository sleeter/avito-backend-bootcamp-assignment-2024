package handler

import (
	"backend-bootcamp-assignment-2024/internal/model/dto/request"
	"backend-bootcamp-assignment-2024/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetHouse(ctx *gin.Context, service *service.Service) error {
	houseId, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	//TODO: add middleware and param to ctx, get isModerator from ctx
	isModerator := false
	err = validateGetHouseFields(houseId)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	resp, err := service.FlatService.GetFlats(ctx, int32(houseId), isModerator)
	if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, gin.H{
		"flats": resp,
	})
	return nil
}
func validateGetHouseFields(houseId int64) error {
	if houseId <= 0 {
		return errors.New("house id must be greater than 0")
	}
	return nil
}

func CreateHouse(ctx *gin.Context, service *service.Service) error {
	var req request.House
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	err := validateCreateHouseField(req)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	resp, err := service.HouseService.CreateHouse(ctx, req)
	if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, resp)
	return nil
}
func validateCreateHouseField(req request.House) error {
	if len(req.Address) == 0 {
		return errors.New("address must not be empty")
	}
	if req.Year < 0 {
		return errors.New("year must not be less than 0")
	}
	if req.Developer != nil && len(*req.Developer) == 0 {
		return errors.New("developer must not be empty")
	}
	return nil
}

func SubscribeHouse(ctx *gin.Context, service *service.Service) error {
	return nil
}
