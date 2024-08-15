package usecases

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/entities"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/repositories"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"
)

type TransactionUseCase interface {
	Deposit(ctx context.Context, transaction entities.Transaction) (*models.Transaction, error)
	WithDraw(ctx context.Context, transaction entities.Transaction) (*models.Transaction, error)
	GetTransactionsByAccountId(ctx context.Context, accountId string) ([]*models.Transaction, error)
	GetTransactionByCustomerIdAndOutLimit(ctx context.Context, customerId string, upperLimit float64) ([]*models.Customer, error)
}

type transactionUseCase struct {
	logger                *logrus.Logger
	accountRepository     repositories.AccountRepository
	customerRepository    repositories.CustomerRepository
	transactionRepository repositories.TransactionRepository
}

func NewTransactionUseCase(
	logger *logrus.Logger,
	accountRepository repositories.AccountRepository,
	customerRepository repositories.CustomerRepository,
	transactionRepository repositories.TransactionRepository,
) TransactionUseCase {
	return &transactionUseCase{
		logger:                logger,
		accountRepository:     accountRepository,
		customerRepository:    customerRepository,
		transactionRepository: transactionRepository,
	}
}

func (tuc *transactionUseCase) Deposit(ctx context.Context, transaction entities.Transaction) (*models.Transaction, error) {
	return nil, nil
}

func (tuc *transactionUseCase) WithDraw(ctx context.Context, transaction entities.Transaction) (*models.Transaction, error) {
	return nil, nil
}

func (tuc *transactionUseCase) GetTransactionsByAccountId(ctx context.Context, accountId string) ([]*models.Transaction, error) {
	return nil, nil
}

func (tuc *transactionUseCase) GetTransactionByCustomerIdAndOutLimit(ctx context.Context, customerId string, upperLimit float64) ([]*models.Customer, error) {
	return nil, nil
}
