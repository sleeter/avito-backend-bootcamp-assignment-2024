package repository

import (
	"context"
	"errors"

	"backend-bootcamp-assignment-2024/internal/model/dto/request"
	"backend-bootcamp-assignment-2024/internal/model/entity"

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
	err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Type)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, id string, user request.Register) (*entity.User, error) {
	q := sq.Insert("users").
		Columns("id", "email", "password", "type").
		Values(id, user.Email, user.Password, user.UserType).
		Suffix("RETURNING *").PlaceholderFormat(sq.Dollar)
	res, err := r.executeQuery(ctx, q)
	if err != nil {
		return nil, err
	}
	return &res[0], nil
}

func (r *UserRepository) GetById(ctx context.Context, id string) (*entity.User, error) {
	q := sq.Select("*").From("users").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)
	users, err := r.executeQuery(ctx, q)
	if err != nil {
		return nil, err
	}
	if len(users) != 1 {
		return nil, errors.New("something went wrong with get user by id")
	}
	return &users[0], nil
}
