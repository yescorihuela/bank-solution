package repositories

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/entities"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/repositories"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"
)

type TransactionRepositoryPostgresql struct {
	db     *pgxpool.Pool
	logger *logrus.Logger
}

func NewTransactionRepositoryPostgresql(
	db *pgxpool.Pool,
	logger *logrus.Logger,
) repositories.TransactionRepository {
	return &TransactionRepositoryPostgresql{
		db:     db,
		logger: logger,
	}
}

func (trp *TransactionRepositoryPostgresql) Deposit(ctx context.Context, transaction entities.Transaction) (*models.Transaction, error) {
	return nil, nil
}

func (trp *TransactionRepositoryPostgresql) WithDraw(ctx context.Context, transaction entities.Transaction) (*models.Transaction, error) {
	return nil, nil
}

func (trp *TransactionRepositoryPostgresql) GetTransactionsByAccountId(ctx context.Context, accountId string) ([]*models.Transaction, error) {
	return nil, nil
}

func (trp *TransactionRepositoryPostgresql) GetTransactionByCustomerIdAndOutLimit(ctx context.Context, customerId string, upperLimit float64) ([]*models.Customer, error) {
	return nil, nil
}
