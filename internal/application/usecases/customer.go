package usecases

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/entities"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/repositories"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"
)

type CustomerUseCase interface {
	Insert(ctx context.Context, customer *entities.Customer) (*models.Customer, error)
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

func (cuc *customerUseCase) Insert(ctx context.Context, customer *entities.Customer) (*models.Customer, error) {
	cuc.logger.Info("Starting customerUseCase.Insert method")
	customerModel, err := cuc.customerRepository.Insert(ctx, customer)
	if err != nil {
		cuc.logger.Errorf("Error during access to customerRepository in customerUseCase.Insert method %s", err)
		return nil, err
	}
	cuc.logger.Info("customerUseCase.Insert executed successfully")
	return customerModel, nil
}

func (cuc *customerUseCase) GetById(ctx context.Context, customerId string) (*models.Customer, error) {
	cuc.logger.Info("Starting customerUseCase.GetById method")
	customerModel, err := cuc.customerRepository.GetById(ctx, customerId)
	if err != nil {
		cuc.logger.Errorf("Error during access to customerRepository in customerUseCase.GetById method %s", err)
		return nil, err
	}
	cuc.logger.Info("customerUseCase.GetById executed successfully")
	return customerModel, nil
}
