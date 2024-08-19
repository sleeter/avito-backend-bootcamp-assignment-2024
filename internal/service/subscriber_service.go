//go:generate mockgen -source ./subscriber_service.go -destination=./mocks/subscriber_service.go -package=mock_service

package service

import (
	"context"
	"fmt"

	"backend-bootcamp-assignment-2024/internal/model/dto/request"
	"backend-bootcamp-assignment-2024/internal/model/entity"
)

type SubscriberRepository interface {
	GetSubsByHouseId(ctx context.Context, houseId int32) ([]entity.Subscriber, error)
	AddSub(ctx context.Context, req request.Subscriber) error
}

type Sender interface {
	SendEmail(ctx context.Context, recipient string, message string) error
}

type SubscriberService struct {
	Repository         SubscriberRepository
	TransactionManager TransactionManager
	Sender             Sender
}

func NewSubscriberService(r SubscriberRepository, manager TransactionManager, sender Sender) SubscriberService {
	return SubscriberService{
		Repository:         r,
		TransactionManager: manager,
		Sender:             sender,
	}
}

func (s *SubscriberService) CreateSubscriber(ctx context.Context, req request.Subscriber) error {
	return s.Repository.AddSub(ctx, req)
}

func (s *SubscriberService) GetSubsByHouseId(ctx context.Context, houseId int32) ([]entity.Subscriber, error) {
	return s.Repository.GetSubsByHouseId(ctx, houseId)
}

func (s *SubscriberService) SendMessageToSubscribers(ctx context.Context, houseId int32) error {
	subs, err := s.GetSubsByHouseId(ctx, houseId)
	if err != nil {
		return err
	}
	for _, sub := range subs {
		err := s.Sender.SendEmail(ctx, sub.Email, fmt.Sprintf("new flat has appeared in the house with id=%d", houseId))
		if err != nil {
			return err
		}
	}
	return nil
}
