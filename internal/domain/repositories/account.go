package repositories

import (
	"context"

	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/entities"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"
)

type AccountRepository interface {
	Insert(ctx context.Context, account *entities.Account) (*models.Account, error)
	GetById(ctx context.Context, customerId, accountId string) (*models.Account, error)
	GetAccountWithTransactionsByAccountId(ctx context.Context, lastTransactions int, customerId, accountId string) (*models.Account, error)
	GetAccountWithTransactionsByAccountIdAndMonth(ctx context.Context, month, year int, customerId, accountId string) (*models.Account, error)
}
