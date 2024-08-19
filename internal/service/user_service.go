//go:generate mockgen -source ./user_service.go -destination=./mocks/user_service.go -package=mock_service

package service

import (
	"context"

	"backend-bootcamp-assignment-2024/internal/mapper"
	"backend-bootcamp-assignment-2024/internal/model/dto/request"
	"backend-bootcamp-assignment-2024/internal/model/dto/response"
	"backend-bootcamp-assignment-2024/internal/model/entity"
	"backend-bootcamp-assignment-2024/internal/pkg/auth"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	CreateUser(ctx context.Context, id string, user request.Register) (*entity.User, error)
	GetById(ctx context.Context, id string) (*entity.User, error)
}

type UserService struct {
	Repository         UserRepository
	TransactionManager TransactionManager
}

func NewUserService(r UserRepository, manager TransactionManager) UserService {
	return UserService{
		Repository:         r,
		TransactionManager: manager,
	}
}

func (s *UserService) Register(ctx context.Context, req request.Register) (*response.Register, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	req.Password = string(passHash)
	if err != nil {
		return nil, err
	}
	id := uuid.New().String()
	res, err := s.Repository.CreateUser(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return mapper.UserEntityToRegisterResponse(res), nil
}

func (s *UserService) Login(ctx context.Context, req request.Login) (*response.Login, error) {
	res, err := s.Repository.GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(req.Password)); err != nil {
		return nil, err
	}
	token, err := auth.CreateUserJWT(req, res.Type)
	if err != nil {
		return nil, err
	}
	return mapper.TokenToResponseLogin(token), nil
}

func (s *UserService) DummyLogin(ctx context.Context, req request.DummyLogin) (*response.Login, error) {
	token, err := auth.CreateDummyJWT(req.UserType)
	if err != nil {
		return nil, err
	}
	return mapper.TokenToResponseLogin(token), nil
}
