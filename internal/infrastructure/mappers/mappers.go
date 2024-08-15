package mappers

import (
	"time"

	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/constants"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/entities"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/shared"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/http/api/requests"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/http/api/responses"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"
)

func FromAccountModelToEntity(account models.Account) entities.Account {
	return entities.Account{}
}

func FromCustomerModelToEntity(customer models.Customer) entities.Customer {
	return entities.Customer{
		Id:        customer.Id,
		Name:      customer.Name,
		Kind:      customer.Kind,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}
}

func FromTransactionModelToEntity(transaction models.Transaction) entities.Transaction {
	return entities.Transaction{}
}

func FromCustomerRequestToEntity(customer requests.CustomerRequest) *entities.Customer {
	id := shared.GenerateNanoId()
	now := time.Now().UTC()
	return &entities.Customer{
		Id:        id,
		Name:      customer.Name,
		Kind:      customer.Kind,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func FromAccountRequestToEntity(account requests.AccountRequest, customerId string) *entities.Account {
	id := shared.GenerateUuid()
	now := time.Now().UTC()
	return &entities.Account{
		Id:         id,
		Kind:       account.Kind,
		CustomerId: customerId,
		Balance:    account.Balance,
		City:       account.City,
		Country:    account.Country,
		Currency:   account.Currency,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

func FromTransactionRequestToEntity(transaction requests.TransactionRequest, customerId, accountId string) *entities.Transaction {
	id := shared.GenerateUlid()
	now := time.Now().UTC()
	return &entities.Transaction{
		Id:        id,
		AccountId: accountId,
		Kind:      transaction.Kind,
		Balance:   transaction.Amount,
		Status:    constants.Pending,
		CreatedAt: now,
	}
}

func FromCustomerModelToResponse(customer *models.Customer) responses.Customer {
	return responses.Customer{
		Id:        customer.Id,
		Name:      customer.Name,
		Kind:      customer.Kind,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}
}

func FromAccountModelToResponse(account *models.Account) responses.Account {
	return responses.Account{
		Id:         account.Id,
		CustomerId: account.CustomerId,
		Kind:       account.Kind,
		Balance:    account.Balance,
		City:       account.City,
		Country:    account.Country,
		Currency:   account.Currency,
		CreatedAt:  account.CreatedAt,
		UpdatedAt:  account.UpdatedAt,
	}
}
