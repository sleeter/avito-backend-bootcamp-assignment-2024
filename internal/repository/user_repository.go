package repository

import (
	"backend-bootcamp-assignment-2024/internal/model/dto/request"
	"backend-bootcamp-assignment-2024/internal/model/entity"
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

type UserRepository struct {
	QueryManager QueryManager
}

func NewUserRepository(manager QueryManager) *UserRepository {
	return &UserRepository{QueryManager: manager}
}

func (r *UserRepository) executeQuery(ctx context.Context, query sq.Sqlizer) ([]entity.User, error) {
	rows, err := r.QueryManager.QuerySq(ctx, query)
	if err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]entity.User, 0)
	for rows.Next() {
		user, err := toUser(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, user)
	}
	return result, nil
}

func toUser(rows pgx.Rows) (entity.User, error) {
	var user entity.User
	err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Type, &user.Dummy)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, id string, user request.Register) error {
	q := sq.Insert("users").
		Columns("id", "email", "password", "type").
		Values(id, user.Email, user.Password, user.UserType).
		Suffix("RETURNING id").PlaceholderFormat(sq.Dollar)
	_, err := r.executeQuery(ctx, q)
	return err
}

func (r *UserRepository) GetById(ctx context.Context, id string) (entity.User, error) {
	q := sq.Select("*").From("users").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)
	users, err := r.executeQuery(ctx, q)
	if err != nil {
		return entity.User{}, err
	}
	if len(users) != 1 {
		return entity.User{}, errors.New("something went wrong with get user by id")
	}
	return users[0], nil
}
