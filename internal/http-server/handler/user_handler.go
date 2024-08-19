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

func Register(ctx *gin.Context, service *service.Service) error {
	var req request.Register
	if err := ctx.Bind(&req); err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	err := validateRegisterFields(req)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	resp, err := service.UserService.Register(ctx, req)
	if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, resp)
	return nil
}

func validateRegisterFields(req request.Register) error {
	if len(req.Email) == 0 {
		return errors.New("email is required")
	}
	if len(req.Password) == 0 {
		return errors.New("password is required")
	}
	if !slices.Contains(entity.UserTypes, req.UserType) {
		return errors.New("invalid user type")
	}
	return nil
}

func Login(ctx *gin.Context, service *service.Service) error {
	var req request.Login
	if err := ctx.Bind(&req); err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	err := validateLoginFields(req)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	resp, err := service.UserService.Login(ctx, req)
	if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, resp)
	return nil
}

func validateLoginFields(req request.Login) error {
	if len(req.Id) == 0 {
		return errors.New("id is required")
	}
	if len(req.Password) == 0 {
		return errors.New("password is required")
	}
	return nil
}

func DummyLogin(ctx *gin.Context, service *service.Service) error {
	var req request.DummyLogin
	req.UserType = ctx.Query("user_type")
	err := validateDummyLoginFields(req)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	resp, err := service.UserService.DummyLogin(ctx, req)
	if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, resp)
	return nil
}

func validateDummyLoginFields(req request.DummyLogin) error {
	if !slices.Contains(entity.UserTypes, req.UserType) {
		return errors.New("invalid user type")
	}
	return nil
}
