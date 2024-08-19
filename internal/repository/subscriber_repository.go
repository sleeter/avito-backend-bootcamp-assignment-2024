package repository

import (
	"backend-bootcamp-assignment-2024/internal/model/dto/request"
	"backend-bootcamp-assignment-2024/internal/model/entity"
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

type SubscriberRepository struct {
	QueryManager QueryManager
}

func NewSubscriberRepository(manager QueryManager) *SubscriberRepository {
	return &SubscriberRepository{QueryManager: manager}
}

func (r *SubscriberRepository) executeQuery(ctx context.Context, query sq.Sqlizer) ([]entity.Subscriber, error) {
	rows, err := r.QueryManager.QuerySq(ctx, query)
	if err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]entity.Subscriber, 0)
	for rows.Next() {
		sub, err := toSubs(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, sub)
	}
	return result, nil
}
func toSubs(rows pgx.Rows) (entity.Subscriber, error) {
	var sub entity.Subscriber
	err := rows.Scan(&sub.Id, &sub.HouseId, &sub.Email)
	if err != nil {
		return entity.Subscriber{}, err
	}
	return sub, nil
}

func (r *SubscriberRepository) GetSubsByHouseId(ctx context.Context, houseId int32) ([]entity.Subscriber, error) {
	q := sq.Select("*").
		From("subscribers").
		Where(sq.Eq{"house_id": houseId}).
		PlaceholderFormat(sq.Dollar)
	res, err := r.executeQuery(ctx, q)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *SubscriberRepository) AddSub(ctx context.Context, req request.Subscriber) error {
	q := sq.Insert("subscribers").
		Columns("house_id", "email").
		Values(req.HouseId, req.Email).
		PlaceholderFormat(sq.Dollar)
	_, err := r.executeQuery(ctx, q)
	if err != nil {
		return err
	}
	return nil
}
