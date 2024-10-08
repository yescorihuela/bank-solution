package usecases

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/entities"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/repositories"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"
)

type AccountUseCase interface {
	Insert(ctx context.Context, account *entities.Account) (*models.Account, error)
	GetById(ctx context.Context, customerId, accountId string) (*models.Account, error)
	GetLastTransactionsById(ctx context.Context, lastTransactions int, customerId, accountId string) (*models.Account, error)
	GetLastTransactionsByAccountIdAndMonth(ctx context.Context, month, year int, customerId, accountId string) (*models.Account, error)
}

type accountUseCase struct {
	logger                *logrus.Logger
	accountRepository     repositories.AccountRepository
	customerRepository    repositories.CustomerRepository
	transactionRepository repositories.TransactionRepository
}

func NewAccountUseCase(
	logger *logrus.Logger,
	accountRepository repositories.AccountRepository,
	customerRepository repositories.CustomerRepository,
	transactionRepository repositories.TransactionRepository,
) AccountUseCase {
	return &accountUseCase{
		logger:                logger,
		accountRepository:     accountRepository,
		customerRepository:    customerRepository,
		transactionRepository: transactionRepository,
	}
}

func (auc *accountUseCase) Insert(ctx context.Context, account *entities.Account) (*models.Account, error) {
	auc.logger.Info("Starting accountUseCase.Insert method")
	accountModel, err := auc.accountRepository.Insert(ctx, account)
	if err != nil {
		auc.logger.Errorf("Error during access to accountRepository in accountUseCase.Insert method %s", err)
		return nil, err
	}
	auc.logger.Info("accountUseCase.Insert executed successfully")
	return accountModel, nil
}

func (auc *accountUseCase) GetById(ctx context.Context, customerId, accountId string) (*models.Account, error) {
	auc.logger.Info("Starting accountUseCase.GetById method")
	accountModel, err := auc.accountRepository.GetById(ctx, customerId, accountId)
	if err != nil {
		auc.logger.Errorf("Error during access to accountRepository in accountUseCase.GetById method %s", err)
		return nil, err
	}
	auc.logger.Info("accountUseCase.GetById executed successfully")
	return accountModel, nil
}

func (auc *accountUseCase) GetLastTransactionsById(ctx context.Context, lastTransactions int, customerId, accountId string) (*models.Account, error) {
	auc.logger.Info("Starting accountUseCase.GetLastTransactionsById method")
	accountModel, err := auc.accountRepository.GetAccountWithTransactionsByAccountId(ctx, lastTransactions, customerId, accountId)
	if err != nil {
		auc.logger.Errorf("Error during access to accountRepository in accountUseCase.GetLastTransactionsById method %s", err)
		return nil, err
	}
	auc.logger.Info("accountUseCase.GetLastTransactionsById executed successfully")
	return accountModel, nil
}

func (auc *accountUseCase) GetLastTransactionsByAccountIdAndMonth(ctx context.Context, month, year int, customerId, accountId string) (*models.Account, error) {
	auc.logger.Info("Starting accountUseCase.GetLastTransactionsByAccountIdAndMonth method")
	accountModel, err := auc.accountRepository.GetAccountWithTransactionsByAccountIdAndMonth(ctx, month, year, customerId, accountId)
	if err != nil {
		auc.logger.Errorf("Error during access to accountRepository in accountUseCase.GetLastTransactionsByAccountIdAndMonth method %s", err)
		return nil, err
	}
	auc.logger.Info("accountUseCase.GetLastTransactionsByAccountIdAndMonth executed successfully")
	return accountModel, nil
}
