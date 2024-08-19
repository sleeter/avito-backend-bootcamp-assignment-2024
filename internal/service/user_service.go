package service

type UserRepository interface{}

type UserService struct {
	Repository         HouseRepository
	TransactionManager TransactionManager
}

func NewUserService(r UserRepository, manager TransactionManager) UserService {
	return UserService{
		Repository:         r,
		TransactionManager: manager,
	}
}
