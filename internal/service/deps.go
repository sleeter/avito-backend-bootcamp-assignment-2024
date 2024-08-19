//go:generate mockgen -source ./deps.go -destination=./mocks/service.go -package=mock_service

package service

import (
	"context"

	"backend-bootcamp-assignment-2024/internal/pkg/pgdb"
)

type TransactionManager interface {
	ReadonlyTx(ctx context.Context, callback pgdb.TransactionCallback) error
	Tx(ctx context.Context, callback pgdb.TransactionCallback) error
}
