//go:generate mockgen -source ./deps.go -destination=./mocks/repository.go -package=mock_repository

package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

type QueryManager interface {
	QuerySq(ctx context.Context, query sq.Sqlizer) (pgx.Rows, error)
}
