package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/entities"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/repositories"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"
)

type TransactionRepositoryPostgresql struct {
	db *pgx.Conn
}

func NewTransactionRepositoryPostgresql(db *pgx.Conn) repositories.TransactionRepository {
	return &TransactionRepositoryPostgresql{db: db}
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
