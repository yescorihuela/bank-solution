package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/entities"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/repositories"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"
)

type AccountRepositoryPostgresql struct {
	db *pgx.Conn
}

func NewAccountRepositoryPostgresql(db *pgx.Conn) repositories.AccountRepository {
	return &AccountRepositoryPostgresql{db: db}
}

func (arp *AccountRepositoryPostgresql) Insert(ctx context.Context, account entities.Account) (*models.Account, error) {
	return nil, nil
}

func (arp *AccountRepositoryPostgresql) GetById(ctx context.Context, accountId string) (*models.Account, error) {
	return nil, nil
}

func (arp *AccountRepositoryPostgresql) GetAccountsByCustomerId(ctx context.Context, customerId string) ([]*models.Account, error) {
	return nil, nil
}
