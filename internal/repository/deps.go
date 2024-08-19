//go:generate mockgen -source ./deps.go -destination=./mocks/repository.go -package=mock_repository

package repository

import (
	"backend-bootcamp-assignment-2024/internal/model/entity"
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

type QueryManager interface {
	QuerySq(ctx context.Context, query sq.Sqlizer) (pgx.Rows, error)
}
type CacheForFlat interface {
	Add(int32, []entity.Flat) bool
	Get(int32) ([]entity.Flat, bool)
	Contains(int32) bool
}
