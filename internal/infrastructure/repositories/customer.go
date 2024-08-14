package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/entities"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/repositories"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"
)

type CustomerRepositoryPostgresql struct {
	db     *pgx.Conn
	logger *logrus.Logger
}

func NewCustomerRepositoryPostgresql(db *pgx.Conn) repositories.CustomerRepository {
	return &CustomerRepositoryPostgresql{db: db}
}

func (crp *CustomerRepositoryPostgresql) Insert(ctx context.Context, customer entities.Customer) (*models.Customer, error) {
	return nil, nil
}

func (crp *CustomerRepositoryPostgresql) GetById(ctx context.Context, customerId string) (*models.Customer, error) {
	return nil, nil
}
