package handler

import (
	"backend-bootcamp-assignment-2024/internal/model/dto/request"
	"backend-bootcamp-assignment-2024/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateFlat(ctx *gin.Context, service *service.Service) error {
	var req request.CreateFlat
	if err := ctx.Bind(&req); err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	//TODO: validate fields
	resp, err := service.FlatService.AddFlat(ctx, req)
	if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, resp)
	return nil
}

func UpdateFlat(ctx *gin.Context, service *service.Service) error {
	var req request.UpdateFlat
	if err := ctx.Bind(&req); err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	//TODO: validate fields
	resp, err := service.FlatService.UpdateFlat(ctx, req)
	if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, resp)
	return nil
}
