package mappers

import (
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/entities"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"
)

func FromAccountModelToEntity(account models.Account) entities.Account {
	return entities.Account{}
}

func FromCustomerModelToEntity(customer models.Customer) entities.Customer {
	return entities.Customer{}
}

func FromTransactionModelToEntity(transaction models.Transaction) entities.Transaction {
	return entities.Transaction{}
}
