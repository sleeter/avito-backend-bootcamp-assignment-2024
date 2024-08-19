package handler

import (
	"backend-bootcamp-assignment-2024/internal/model/dto/request"
	"backend-bootcamp-assignment-2024/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetHouse(ctx *gin.Context, service *service.Service) error {
	HouseId, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
	//TODO: add middleware and param to ctx, get isModerator from ctx
	isModerator := false
	//TODO: validate fields
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	resp, err := service.FlatService.GetFlats(ctx, int32(HouseId), isModerator)
	if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, gin.H{
		"flats": resp,
	})
	return nil
}

func CreateHouse(ctx *gin.Context, service *service.Service) error {
	var req request.House
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	//TODO: validate fields
	resp, err := service.HouseService.CreateHouse(ctx, req)
	if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, resp)
	return nil
}

func SubscribeHouse(ctx *gin.Context, service *service.Service) error {
	return nil
}
