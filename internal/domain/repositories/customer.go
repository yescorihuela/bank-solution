package repositories

import (
	"context"

	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/entities"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"
)

type CustomerRepository interface {
	Insert(ctx context.Context, customer entities.Customer) (*models.Customer, error)
	GetById(ctx context.Context, customerId string) (*models.Customer, error)
}
