package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/constants"
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
			customer_id,
			kind,
			balance,
			city,
			country,
			currency,
			created_at,
			updated_at
		)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING 
		id,
		customer_id,
		kind,
		balance,
		city,
		country,
		currency,
		created_at,
		updated_at
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
	id = $1 AND customer_id = $2
	`)

var getTransactionsByAccountId = shared.Compact(`
		SELECT 
			t.id as transaction_id, t.amount, t.kind, t.status, t.city, t.created_at 
		FROM 
			accounts a 
		INNER JOIN
		transactions t 
			ON(a.id = t.account_id)
		WHERE t.account_id = $1
		ORDER BY t.created_at DESC LIMIT $2
	`)

var getTransactionByAccountIdAndMonth = shared.Compact(`
		SELECT 
			t.id as transaction_id, t.amount, t.kind, t.status, t.city, t.created_at 
		FROM 
			accounts a 
		INNER JOIN
		transactions t 
			ON(a.id = t.account_id)
		WHERE t.account_id = $1 AND 
		EXTRACT(month from t.created_at::date) = $2 AND
		EXTRACT(year from t.created_at::date) = $3
		ORDER BY t.created_at DESC
`)

var checkValidCurrent = shared.Compact(`
	SELECT kind from customers WHERE id = $1
`)

func (arp *AccountRepositoryPostgresql) Insert(ctx context.Context, account *entities.Account) (*models.Account, error) {
	accountModel := models.NewAccountModel()
	var kind int

	err := arp.db.QueryRow(ctx, checkValidCurrent, account.CustomerId).Scan(&kind)
	if err != nil {
		return nil, err
	}

	if kind == constants.Individual && account.Kind == constants.CurrentAccount {
		return nil, fmt.Errorf("the current accounts only work for organizations")
	} else if kind == constants.Organization && account.Kind == constants.SavingAccout {
		return nil, fmt.Errorf("the saving accounts only work for individuals")
	}

	err = arp.db.
		QueryRow(ctx, insertAccountQuery, account.Id, account.CustomerId, account.Kind, account.Balance, account.City, account.Country, account.Currency, account.CreatedAt, account.UpdatedAt).
		Scan(
			&accountModel.Id,
			&accountModel.CustomerId,
			&accountModel.Kind,
			&accountModel.Balance,
			&accountModel.City,
			&accountModel.Country,
			&accountModel.Currency,
			&accountModel.CreatedAt,
			&accountModel.UpdatedAt,
		)

	if err != nil {
		return nil, err
	}
	return &accountModel, nil
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

func (arp *AccountRepositoryPostgresql) GetAccountWithTransactionsByAccountId(ctx context.Context, lastTransactions int, customerId, accountId string) (*models.Account, error) {
	arp.logger.Info("Starting AccountRepositoryPostgresql.GetAccountWithTransactionsByAccountId method")
	accountModel := models.NewAccountModel()
	transactionModels := make([]*models.Transaction, 0)
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
		arp.logger.Errorf("Failing AccountRepositoryPostgresql.GetAccountWithTransactionsByAccountId method %s", err)
		return nil, err
	}
	resultSet, err := arp.db.Query(ctx, getTransactionsByAccountId, accountId, lastTransactions)
	if err != nil {
		arp.logger.Errorf("Failing AccountRepositoryPostgresql.GetAccountWithTransactionsByAccountId method querying transactions %s", err)
		return nil, err
	}

	for resultSet.Next() {
		t := models.NewTransaction()
		err := resultSet.Scan(
			&t.Id,
			&t.Amount,
			&t.Kind,
			&t.Status,
			&t.City,
			&t.CreatedAt,
		)
		if err != nil {
			arp.logger.Errorf("Failing AccountRepositoryPostgresql.GetAccountWithTransactionsByAccountId method fetching transactions rows %s", err)
			return &accountModel, err
		}
		transactionModels = append(transactionModels, &t)
	}
	accountModel.Transactions = transactionModels

	return &accountModel, nil
}

func (arp *AccountRepositoryPostgresql) GetAccountWithTransactionsByAccountIdAndMonth(ctx context.Context, month, year int, customerId, accountId string) (*models.Account, error) {

	arp.logger.Info("Starting AccountRepositoryPostgresql.GetAccountWithTransactionsByAccountIdAndMonth method")
	accountModel := models.NewAccountModel()
	transactionModels := make([]*models.Transaction, 0)

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
		arp.logger.Errorf("Failing AccountRepositoryPostgresql.GetAccountWithTransactionsByAccountIdAndMonth method %s", err)
		return nil, err
	}

	resultSet, err := arp.db.Query(ctx, getTransactionByAccountIdAndMonth, accountId, month, year)
	if err != nil {
		arp.logger.Errorf("Failing AccountRepositoryPostgresql.GetAccountWithTransactionsByAccountIdAndMonth method querying transactions %s", err)
		return nil, err
	}

	for resultSet.Next() {
		t := models.NewTransaction()
		err := resultSet.Scan(
			&t.Id,
			&t.Amount,
			&t.Kind,
			&t.Status,
			&t.City,
			&t.CreatedAt,
		)
		if err != nil {
			arp.logger.Errorf("Failing AccountRepositoryPostgresql.GetAccountWithTransactionsByAccountIdAndMonth method fetching transactions rows %s", err)
			return &accountModel, err
		}
		transactionModels = append(transactionModels, &t)
	}
	accountModel.Transactions = transactionModels

	return &accountModel, nil
}
