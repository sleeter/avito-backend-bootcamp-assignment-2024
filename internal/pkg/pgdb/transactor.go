package pgdb

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type TransactionManager struct {
	pool *pgxpool.Pool
}

func NewTransactionManager(pool *pgxpool.Pool) *TransactionManager {
	return &TransactionManager{pool: pool}
}

// TransactionCallback represents function that will be executed withing single db transaction
type TransactionCallback func(ctx context.Context) error

// ReadonlyTx makes transaction reading without commit.
func (tm *TransactionManager) ReadonlyTx(ctx context.Context, callback TransactionCallback) error {
	tx, err := tm.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	ctx = context.WithValue(ctx, txCtxKey{}, tx)

	if err := callback(ctx); err != nil {
		return err
	}

	return nil
}

// Tx executes a callback within a single transaction.
func (tm *TransactionManager) Tx(ctx context.Context, callback TransactionCallback) error {
	tx, err := tm.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	ctx = context.WithValue(ctx, txCtxKey{}, tx)

	if err = callback(ctx); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
