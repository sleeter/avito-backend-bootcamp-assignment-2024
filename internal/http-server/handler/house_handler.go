package handler

import (
	"backend-bootcamp-assignment-2024/internal/model/dto/request"
	"backend-bootcamp-assignment-2024/internal/model/entity"
	"backend-bootcamp-assignment-2024/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetHouse(ctx *gin.Context, service *service.Service) error {
	log.SetPrefix("http-server.handler.GetHouse")
	houseId, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
	if err != nil {
		log.Println(err.Error())
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	userType := ctx.GetString("User-Type")
	isModerator, err := validateGetHouseFields(houseId, userType)
	if err != nil {
		log.Println(err.Error())
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	resp, err := service.FlatService.GetFlats(ctx, int32(houseId), isModerator)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	ctx.JSON(http.StatusOK, gin.H{
		"flats": resp,
	})
	return nil
}
func validateGetHouseFields(houseId int64, userType string) (bool, error) {
	if houseId <= 0 {
		return false, errors.New("house id must be greater than 0")
	}
	if len(userType) == 0 {
		return false, errors.New("user type can not be empty")
	} else if userType == entity.USERTYPE_MODERATOR {
		return true, nil
	}
	return false, nil
}

func CreateHouse(ctx *gin.Context, service *service.Service) error {
	log.SetPrefix("http-server.handler.CreateHouse")
	var req request.House
	if err := ctx.Bind(&req); err != nil {
		log.Println(err.Error())
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	err := validateCreateHouseField(req)
	if err != nil {
		log.Println(err.Error())
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	resp, err := service.HouseService.CreateHouse(ctx, req)
	if err != nil {
		log.Println(err.Error())
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
	log.SetPrefix("http-server.handler.SubscribeHouse")
	houseId, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
	if err != nil {
		log.Println(err.Error())
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	var req request.Subscriber
	if err := ctx.Bind(&req); err != nil {
		log.Println(err.Error())
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	req.HouseId = int32(houseId)
	//TODO: validate email
	err = validateSubscribeHouseFields(req)
	if err != nil {
		log.Println(err.Error())
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	err = service.SubscriberService.CreateSubscriber(ctx, req)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func validateSubscribeHouseFields(req request.Subscriber) error {
	if req.HouseId <= 0 {
		return errors.New("house id must be greater than 0")
	}
	return nil
}
