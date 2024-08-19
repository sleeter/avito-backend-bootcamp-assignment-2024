package service

type FlatRepository interface{}

type FlatService struct {
	Repository         HouseRepository
	TransactionManager TransactionManager
}

func NewFlatService(r FlatRepository, manager TransactionManager) FlatService {
	return FlatService{
		Repository:         r,
		TransactionManager: manager,
	}
}
