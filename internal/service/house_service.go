package service

type HouseRepository interface{}

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
