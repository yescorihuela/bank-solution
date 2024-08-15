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

type AccountRepositoryPostgresql struct {
	db     *pgxpool.Pool
	logger *logrus.Logger
}

func NewAccountRepositoryPostgresql(
	db *pgxpool.Pool,
	logger *logrus.Logger,
) repositories.AccountRepository {
	return &AccountRepositoryPostgresql{
		db:     db,
		logger: logger,
	}
}

var insertAccountQuery = shared.Compact(`
	INSERT INTO 
		accounts(
			id,
			kind,
			customer_id,
			balance,
			city,
			country,
			currency,
			created_at,
			updated_at
		)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING *	
`)

var getAccountByIdQuery = shared.Compact(`
	SELECT id,
			kind,
			balance,
			city,
			country,
			currency,
			created_at,
			updated_at
	FROM accounts WHERE
	id = $1 AND customer_id = $2`)

func (arp *AccountRepositoryPostgresql) Insert(ctx context.Context, account *entities.Account) error {
	_, err := arp.db.Exec(ctx, insertAccountQuery, account.Id, account.Kind, account.CustomerId, account.Balance, account.City, account.Country, account.Currency, account.CreatedAt, account.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (arp *AccountRepositoryPostgresql) GetById(ctx context.Context, customerId, accountId string) (*models.Account, error) {
	arp.logger.Info("Starting AccountRepositoryPostgresql.GetById method")
	accountModel := models.NewAccountModel()
	err := arp.db.QueryRow(ctx, getAccountByIdQuery, accountId, customerId).Scan(
		&accountModel.Id,
		&accountModel.Kind,
		&accountModel.Balance,
		&accountModel.City,
		&accountModel.Country,
		&accountModel.Currency,
		&accountModel.CreatedAt,
		&accountModel.UpdatedAt,
	)
	if err != nil {
		arp.logger.Error("Failing AccountRepositoryPostgresql.GetById method")
		return nil, err
	}
	arp.logger.Info("AccountRepositoryPostgresql.GetById method Finished")
	return &accountModel, nil
}

func (arp *AccountRepositoryPostgresql) GetAccountsByCustomerId(ctx context.Context, customerId string) ([]*models.Account, error) {
	return nil, nil
}
