package service

import (
	"backend-bootcamp-assignment-2024/internal/pkg/pgdb"
	"context"
)

type TransactionManager interface {
	ReadonlyTx(ctx context.Context, callback pgdb.TransactionCallback) error
	Tx(ctx context.Context, callback pgdb.TransactionCallback) error
}
