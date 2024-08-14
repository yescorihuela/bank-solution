package repositories

import (
	"context"

	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/entities"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"
)

type TransactionRepository interface {
	Deposit(ctx context.Context, transaction entities.Transaction) (*models.Transaction, error)
	WithDraw(ctx context.Context, transaction entities.Transaction) (*models.Transaction, error)
	GetTransactionsByAccountId(ctx context.Context, accountId string) ([]*models.Transaction, error)
	GetTransactionByCustomerIdAndOutLimit(ctx context.Context, customerId string, upperLimit float64) ([]*models.Customer, error)
}
