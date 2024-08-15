package repositories

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/entities"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/repositories"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/shared"
)

type CustomerRepositoryPostgresql struct {
	db     *pgxpool.Pool
	logger *logrus.Logger
}

func NewCustomerRepositoryPostgresql(
	db *pgxpool.Pool,
	logger *logrus.Logger,
) repositories.CustomerRepository {
	return &CustomerRepositoryPostgresql{
		db:     db,
		logger: logger,
	}
}

var insertCustomerQuery = shared.Compact(`
	INSERT INTO 
	customers 
		(id, name, kind, created_at, updated_at)
	VALUES 
		($1, $2, $3, $4, $5) 
	RETURNING id, name, kind, created_at, updated_at
`)

var getCustomerByIdQuery = shared.Compact(`
	SELECT 
		id, name, kind, created_at, updated_at 
	FROM 
		customers
	WHERE id = $1
`)

func (crp *CustomerRepositoryPostgresql) Insert(ctx context.Context, customer *entities.Customer) (*models.Customer, error) {
	crp.logger.Info("Starting CustomerRepositoryPostgresql.Insert method")
	customerModel := models.NewCustomerModel()
	err := crp.db.QueryRow(ctx, insertCustomerQuery, customer.Id, customer.Name, customer.Kind, customer.CreatedAt, customer.UpdatedAt).
		Scan(
			&customerModel.Id,
			&customerModel.Name,
			&customerModel.Kind,
			&customerModel.CreatedAt,
			&customerModel.UpdatedAt,
		)
	if err != nil {
		crp.logger.Error("Failing CustomerRepositoryPostgresql.Insert method")
		return nil, err
	}
	crp.logger.Info("CustomerRepositoryPostgresql.Insert method Finished")
	return &customerModel, nil
}

func (crp *CustomerRepositoryPostgresql) GetById(ctx context.Context, customerId string) (*models.Customer, error) {
	crp.logger.Info("Starting CustomerRepositoryPostgresql.GetById method")
	customerModel := models.NewCustomerModel()
	err := crp.db.QueryRow(ctx, getCustomerByIdQuery, customerId).Scan(
		&customerModel.Id,
		&customerModel.Name,
		&customerModel.Kind,
		&customerModel.CreatedAt,
		&customerModel.UpdatedAt,
	)
	if err != nil {
		crp.logger.Error("Failing CustomerRepositoryPostgresql.GetById method")
		return nil, err
	}
	crp.logger.Info("CustomerRepositoryPostgresql.GetById method Finished")
	return &customerModel, nil
}
