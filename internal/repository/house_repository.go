package repository

import (
	"context"
	"errors"
	"time"

	"backend-bootcamp-assignment-2024/internal/model/dto/request"
	"backend-bootcamp-assignment-2024/internal/model/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

type HouseRepository struct {
	QueryManager QueryManager
}

func NewHouseRepository(manager QueryManager) *HouseRepository {
	return &HouseRepository{QueryManager: manager}
}

func (r *HouseRepository) executeQuery(ctx context.Context, query sq.Sqlizer) ([]entity.House, error) {
	rows, err := r.QueryManager.QuerySq(ctx, query)
	if err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]entity.House, 0)
	for rows.Next() {
		house, err := toHouse(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, house)
	}
	return result, nil
}

func toHouse(rows pgx.Rows) (entity.House, error) {
	var house entity.House
	err := rows.Scan(&house.Id, &house.Address, &house.Year, &house.Developer, &house.CreatedAt, &house.UpdateAt)
	if err != nil {
		return entity.House{}, err
	}
	return house, nil
}

func (r *HouseRepository) CreateHouse(ctx context.Context, house request.House) (*entity.House, error) {
	t := time.Now()
	q := sq.Insert("houses").
		Columns("address", "year", "developer", "created_at", "update_at").
		Values(house.Address, house.Year, house.Developer, t, t).
		PlaceholderFormat(sq.Dollar).Suffix("RETURNING *")
	houses, err := r.executeQuery(ctx, q)
	if err != nil {
		return nil, err
	}
	if len(houses) != 1 {
		return nil, errors.New("something went wrong with create house")
	}
	return &houses[0], nil
}

func (r *HouseRepository) UpdateHouse(ctx context.Context, houseId int32, updateTime time.Time) error {
	q := sq.Update("houses").
		Set("update_at", updateTime).
		Where(sq.Eq{"id": houseId}).
		PlaceholderFormat(sq.Dollar)
	_, err := r.executeQuery(ctx, q)
	return err
}
