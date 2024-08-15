package usecases

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/entities"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/repositories"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"
)

type CustomerUseCase interface {
	Insert(ctx context.Context, customer *entities.Customer) error
	GetById(ctx context.Context, customerId string) (*models.Customer, error)
}

type customerUseCase struct {
	logger                *logrus.Logger
	accountRepository     repositories.AccountRepository
	customerRepository    repositories.CustomerRepository
	transactionRepository repositories.TransactionRepository
}

func NewCustomerUseCase(
	logger *logrus.Logger,
	accountRepository repositories.AccountRepository,
	customerRepository repositories.CustomerRepository,
	transactionRepository repositories.TransactionRepository,
) CustomerUseCase {
	return &customerUseCase{
		logger:                logger,
		accountRepository:     accountRepository,
		customerRepository:    customerRepository,
		transactionRepository: transactionRepository,
	}
}

func (cuc *customerUseCase) Insert(ctx context.Context, customer *entities.Customer) error {
	cuc.logger.Info("Starting customerUseCase.Insert method")
	err := cuc.customerRepository.Insert(ctx, customer)
	if err != nil {
		cuc.logger.Error("Error during access to customerRepository in customerUseCase.Insert method")
		return err
	}
	cuc.logger.Info("customerUseCase.Insert executed successfully")
	return nil
}

func (cuc *customerUseCase) GetById(ctx context.Context, customerId string) (*models.Customer, error) {
	cuc.logger.Info("Starting customerUseCase.GetById method")
	customerModel, err := cuc.customerRepository.GetById(ctx, customerId)
	if err != nil {
		cuc.logger.Error("Error during access to customerRepository in customerUseCase.GetById method")
		return nil, err
	}
	cuc.logger.Info("customerUseCase.GetById executed successfully")
	return customerModel, nil
}
