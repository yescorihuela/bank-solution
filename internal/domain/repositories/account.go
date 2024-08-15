package repositories

import (
	"context"

	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/entities"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"
)

type AccountRepository interface {
	Insert(ctx context.Context, account *entities.Account) error
	GetById(ctx context.Context, customerId, accountId string) (*models.Account, error)
	GetAccountsByCustomerId(ctx context.Context, customerId string) ([]*models.Account, error)
}
