package usecases

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/entities"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/repositories"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"
)

type AccountUseCase interface {
	Insert(ctx context.Context, account *entities.Account) error
	GetById(ctx context.Context, customerId, accountId string) (*models.Account, error)
	GetAccountsByCustomerId(ctx context.Context, customerId string) ([]*models.Account, error)
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

func (auc *accountUseCase) Insert(ctx context.Context, account *entities.Account) error {
	auc.logger.Info("Starting accountUseCase.Insert method")
	err := auc.accountRepository.Insert(ctx, account)
	if err != nil {
		auc.logger.Error("Error during access to accountRepository in accountUseCase.Insert method")
		return err
	}
	auc.logger.Info("accountUseCase.Insert executed successfully")
	return nil
}

func (auc *accountUseCase) GetById(ctx context.Context, customerId, accountId string) (*models.Account, error) {
	auc.logger.Info("Starting accountUseCase.GetById method")
	accountModel, err := auc.accountRepository.GetById(ctx, customerId, accountId)
	if err != nil {
		auc.logger.Error("Error during access to accountRepository in accountUseCase.GetById method")
		return nil, err
	}
	auc.logger.Info("accountUseCase.GetById executed successfully")
	return accountModel, nil
}

func (auc *accountUseCase) GetAccountsByCustomerId(ctx context.Context, customerId string) ([]*models.Account, error) {
	return nil, nil
}
