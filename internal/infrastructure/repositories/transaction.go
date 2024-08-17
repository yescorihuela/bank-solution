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
	db     postgresql.PGXQueryer
	logger *logrus.Logger
	now    func() time.Time
}

func NewTransactionQueries(
	db *pgx.Tx,
	logger *logrus.Logger,
) *transactionQueries {
	return &transactionQueries{
		db:     *db,
		logger: logger,
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
	SELECT balance FROM accounts WHERE id = $1 AND customer_id = $2
`)

var updateTransactionStatus = shared.Compact(`
	UPDATE transactions SET status = $1 WHERE id = $2 AND account_id = $3
`)

var getProcessedTransaction = shared.Compact(`
	SELECT 
		id, account_id, amount, kind, status, city, created_at 
	FROM transactions
	WHERE id = $1 AND account_id = $2
`)

var updateAccountBalanceForDeposit = shared.Compact(`
	UPDATE 
		accounts 
	SET balance = balance + $1, updated_at = $2
	WHERE 
		id = $3 AND customer_id = $4
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
		id = $3 AND customer_id = $4
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

func (trp *TransactionRepositoryPostgresql) CreateTransaction(ctx context.Context, transaction entities.Transaction, customerId string) (*models.Transaction, error) {
	transactionModel := models.NewTransaction()
	trp.logger.Info("Starting TransactionRepositoryPostgresql.CreateTransaction method")

	err := postgresql.WithTX(ctx, trp.db, func(tx *pgx.Tx) error {
		trp.logger.Info("Starting transaction WithTX method...")

		q := NewTransactionQueries(tx, trp.logger)
		err := q.createTransaction(ctx, transaction, customerId)
		if err != nil {
			trp.logger.Errorf("failing transaction WithTX method... %v", err)
			return err
		}

		trp.logger.Info("Finishing transaction WithTX method...")
		return nil
	})
	if err != nil {
		return nil, err
	}
	err = trp.db.QueryRow(ctx, getProcessedTransaction, transaction.Id, transaction.AccountId).Scan(
		&transactionModel.Id,
		&transactionModel.AccountId,
		&transactionModel.Amount,
		&transactionModel.Kind,
		&transactionModel.Status,
		&transactionModel.City,
		&transactionModel.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &transactionModel, nil
}

func (txq *transactionQueries) createTransaction(ctx context.Context, transaction entities.Transaction, customerId string) error {
	var err error
	if transaction.Kind == constants.WithDrawal {
		txq.logger.Info("Checking balance for withdrawal operation...")
		var balance float64
		err := txq.db.QueryRow(ctx, getAccountCurrentBalance, transaction.AccountId, customerId).Scan(&balance)
		if err != nil {
			txq.logger.Errorf("Withdrawal operation failed %v", err)
			return err
		}

		if transaction.Amount > balance {
			return fmt.Errorf("the amount is greater(%f) than balance(%f)", transaction.Amount, balance)
		}
		txq.logger.Info("Finishing withdrawal operation...")
	}

	txq.logger.Info("Executing transaction record creation...")
	_, err = txq.db.Exec(ctx, createTransaction,
		&transaction.Id,
		&transaction.AccountId,
		&transaction.Amount,
		&transaction.Kind,
		&transaction.Status,
		&transaction.City,
		&transaction.CreatedAt,
	)
	if err != nil {
		return err
	}
	txq.logger.Info("Finishing transaction record creation...")

	query := selectQuery(transaction.Kind)
	txq.logger.Info("Executing account balance update...")
	_, err = txq.db.Exec(ctx, query, transaction.Amount, txq.now().UTC(), transaction.AccountId, customerId)
	if err != nil {
		return err
	}
	txq.logger.Info("Finishing account balance update...")

	txq.logger.Info("Executing transaction status approved update...")
	_, err = txq.db.Exec(ctx, updateTransactionStatus, constants.Approved, transaction.Id, transaction.AccountId)
	if err != nil {
		return err
	}
	txq.logger.Info("Finishing transaction status approved update...")

	return nil
}
