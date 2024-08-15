package usecases

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/entities"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/repositories"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"
)

type TransactionUseCase interface {
	CreateTransaction(ctx context.Context, transaction entities.Transaction) (*models.Transaction, error)
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

func (tuc *transactionUseCase) CreateTransaction(ctx context.Context, transaction entities.Transaction) (*models.Transaction, error) {
	tuc.logger.Info("Starting transactionUseCase.Deposit method")

	transactionModel, err := tuc.transactionRepository.CreateTransaction(ctx, transaction)
	if err != nil {
		tuc.logger.Errorf("Failing transactionUseCase.Deposit method %v", err)
		return nil, err
	}
	tuc.logger.Info("Finishing transactionUseCase.Deposit method")
	return transactionModel, nil
}
