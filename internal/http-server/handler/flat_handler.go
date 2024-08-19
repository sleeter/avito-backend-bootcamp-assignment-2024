package handler

import (
	"backend-bootcamp-assignment-2024/internal/model/dto/request"
	"backend-bootcamp-assignment-2024/internal/model/entity"
	"backend-bootcamp-assignment-2024/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
)

func CreateFlat(ctx *gin.Context, service *service.Service) error {
	var req request.CreateFlat
	if err := ctx.Bind(&req); err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	err := validateCreateFlatFields(req)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	resp, err := service.FlatService.AddFlat(ctx, req)
	if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, resp)
	return nil
}

func validateCreateFlatFields(req request.CreateFlat) error {
	if req.HouseId <= 0 {
		return errors.New("house id must greater than 0")
	}
	if req.Price < 0 {
		return errors.New("price must greater than 0")
	}
	if req.Rooms != nil && *req.Rooms <= 0 {
		return errors.New("rooms must greater or equal to 0")
	}
	return nil
}

func UpdateFlat(ctx *gin.Context, service *service.Service) error {
	var req request.UpdateFlat
	if err := ctx.Bind(&req); err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	err := validateUpdateFlatFields(req)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	resp, err := service.FlatService.UpdateFlat(ctx, req)
	if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, resp)
	return nil
}

func validateUpdateFlatFields(req request.UpdateFlat) error {
	if req.Id <= 0 {
		return errors.New("id must greater than 0")
	}
	if len(req.Status) != 0 && !slices.Contains(entity.FlatStatuses, req.Status) {
		return errors.New("status must be one of enum values")
	}
	return nil
}