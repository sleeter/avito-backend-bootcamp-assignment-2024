package repository

import (
	"backend-bootcamp-assignment-2024/internal/model/dto/request"
	"backend-bootcamp-assignment-2024/internal/model/entity"
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

type FlatRepository struct {
	QueryManager QueryManager
}

func NewFlatRepository(manager QueryManager) *FlatRepository {
	return &FlatRepository{QueryManager: manager}
}

func (r *FlatRepository) executeQuery(ctx context.Context, query sq.Sqlizer) ([]entity.Flat, error) {
	rows, err := r.QueryManager.QuerySq(ctx, query)
	if err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]entity.Flat, 0)
	for rows.Next() {
		flat, err := toFlat(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, flat)
	}
	return result, nil
}

func toFlat(rows pgx.Rows) (entity.Flat, error) {
	var flat entity.Flat
	err := rows.Scan(&flat.Id, &flat.HouseId, &flat.Price, &flat.Rooms, &flat.Status)
	if err != nil {
		return entity.Flat{}, err
	}
	return flat, nil
}

func (r *FlatRepository) GetFlatsByHouseId(ctx context.Context, houseId int32, isModerator bool) ([]entity.Flat, error) {
	q := sq.Select("*").
		From("flats").
		Where(sq.Eq{"house_id": houseId}).
		PlaceholderFormat(sq.Dollar)
	if !isModerator {
		q = q.Where(sq.Eq{"status": entity.FLATSTATUS_APPROVED})
	}
	flats, err := r.executeQuery(ctx, q)
	if err != nil {
		return nil, err
	}
	return flats, nil
}

func (r *FlatRepository) CreateFlat(ctx context.Context, flat request.CreateFlat) (*entity.Flat, error) {
	q := sq.Insert("flats").
		Columns("house_id", "price", "rooms", "status").
		Values(flat.HouseId, flat.Price, *flat.Rooms, flat.Status).
		Suffix("RETURNING *").PlaceholderFormat(sq.Dollar)
	flats, err := r.executeQuery(ctx, q)
	if err != nil {
		return nil, err
	}
	if len(flats) != 1 {
		return nil, errors.New("something went wrong with insert flat")
	}
	return &flats[0], nil
}
func (r *FlatRepository) UpdateFlatStatus(ctx context.Context, flat request.UpdateFlat) (*entity.Flat, error) {
	q := sq.Update("flats").
		Set("status", flat.Status).
		Where(sq.Eq{"id": flat.Id}).
		Suffix("RETURNING *").PlaceholderFormat(sq.Dollar)
	flats, err := r.executeQuery(ctx, q)
	if err != nil {
		return nil, err
	}
	if len(flats) != 1 {
		return nil, errors.New("something went wrong with update flat")
	}
	return &flats[0], nil
}
