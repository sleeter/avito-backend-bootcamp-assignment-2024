package pgdb

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type txCtxKey struct{}

// NewQueryManager returns new *QueryManager.
func NewQueryManager(pool *pgxpool.Pool) *QueryManager {
	return &QueryManager{Pool: pool}
}

type QueryManager struct {
	Pool *pgxpool.Pool
}

// QuerySq executes query with squirrel.
func (qm *QueryManager) QuerySq(ctx context.Context, query sq.Sqlizer) (pgx.Rows, error) {
	tx, withTransaction := transactionFromContext(ctx)

	querySql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	if withTransaction {
		return tx.Query(ctx, querySql, args...)
	}
	return qm.Pool.Query(ctx, querySql, args...)
}

func transactionFromContext(ctx context.Context) (pgx.Tx, bool) {
	if tx := ctx.Value(txCtxKey{}); tx != nil {
		return tx.(pgx.Tx), true
	}
	return nil, false
}
