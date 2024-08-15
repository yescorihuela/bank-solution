package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/constants"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/entities"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/repositories"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/databases/postgresql"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/shared"
)

type TransactionRepositoryPostgresql struct {
	db     *pgxpool.Pool
	logger *logrus.Logger
}

type transactionQueries struct {
	logger *logrus.Logger
	db     postgresql.PGXQueryer
	now    func() time.Time
}

func NewTransactionQueries(
	db *pgxpool.Pool,
	logger *logrus.Logger,
) *transactionQueries {
	return &transactionQueries{
		logger: logger,
		db:     db,
		now:    time.Now,
	}
}

func NewTransactionRepositoryPostgresql(
	db *pgxpool.Pool,
	logger *logrus.Logger,
) repositories.TransactionRepository {
	return &TransactionRepositoryPostgresql{
		db:     db,
		logger: logger,
	}
}

var getAccountCurrentBalance = shared.Compact(`
	SELECT balance FROM accounts WHERE id = $1
`)

var updateAccountBalanceForDeposit = shared.Compact(`
	UPDATE 
		accounts 
	SET balance = balance + $1, updated_at = $2
	WHERE 
		id = $3
	RETURNING id, balance, city, created_at, updated_at
`)

var createTransaction = shared.Compact(`
	INSERT INTO 
		transactions (id, account_id, amount, kind, status, city, created_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id, account_id, amount, kind, status, city, created_at
`)

var updateAccountBalanceForWithDrawal = shared.Compact(`
	UPDATE 
		accounts 
	SET balance = balance - $1, updated_at = $2
	WHERE 
		id = $3
	RETURNING id, balance, city, created_at, updated_at
`)

func selectQuery(kind int) string {
	var query string
	switch kind {
	case constants.Deposit:
		query = updateAccountBalanceForDeposit
	case constants.WithDrawal:
		query = updateAccountBalanceForWithDrawal
	}
	return query
}

func (trp *TransactionRepositoryPostgresql) CreateTransaction(ctx context.Context, transaction entities.Transaction) (*models.Transaction, error) {
	transactionModel := models.NewTransaction()
	trp.logger.Info("Starting TransactionRepositoryPostgresql.Deposit method")

	err := postgresql.WithTX(ctx, trp.db, func(tx pgx.Tx) error {
		trp.logger.Info("Starting deposit transaction...")

		q := NewTransactionQueries(trp.db, trp.logger)
		t, err := q.createTransaction(ctx, transaction)
		transactionModel = *t
		if err != nil {
			trp.logger.Errorf("Failing deposit transaction... %v", err)
			return err
		}
		trp.logger.Info("Finishing deposit transaction...")
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &transactionModel, nil
}

func (txq *transactionQueries) createTransaction(ctx context.Context, transaction entities.Transaction) (*models.Transaction, error) {
	txModel := models.NewTransaction()
	var err error
	if transaction.Kind == constants.WithDrawal {
		var balance float64
		err := txq.db.QueryRow(ctx, getAccountCurrentBalance, transaction.AccountId).Scan(&balance)
		if err != nil {
			return nil, err
		}

		if transaction.Amount > balance {
			return nil, fmt.Errorf("the amount is greater(%f) than balance(%f)", transaction.Amount, balance)
		}
	}

	err = txq.db.QueryRow(ctx, createTransaction,
		&transaction.Id,
		&transaction.AccountId,
		&transaction.Amount,
		&transaction.Kind,
		&transaction.Status,
		&transaction.City,
		&transaction.CreatedAt,
	).Scan(
		&txModel.Id,
		&txModel.AccountId,
		&txModel.Amount,
		&txModel.Kind,
		&txModel.Status,
		&txModel.City,
		&txModel.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	account := models.NewAccountModel()
	err = txq.db.QueryRow(ctx, selectQuery(transaction.Kind), transaction.Amount, txq.now().UTC(), transaction.AccountId).Scan(
		&account.Id,
		&account.Balance,
		&account.City,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &txModel, nil
}
