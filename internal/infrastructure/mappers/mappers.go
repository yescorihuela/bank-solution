package mappers

import (
	"strconv"
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

func FromTransactionRequestToEntity(transaction requests.TransactionRequest, accountId string) *entities.Transaction {
	id := shared.GenerateUlid()
	now := time.Now().UTC()
	kind, _ := strconv.Atoi(transaction.Kind)
	return &entities.Transaction{
		Id:        id,
		AccountId: accountId,
		Amount:    transaction.Amount,
		Kind:      kind,
		Status:    constants.Pending,
		City:      transaction.City,
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

func FromTransactionModelToResponse(transaction *models.Transaction) responses.Transaction {
	return responses.Transaction{
		Id:        transaction.Id,
		AccountId: transaction.AccountId,
		Amount:    transaction.Amount,
		Kind:      transaction.Kind,
		Status:    transaction.Status,
		City:      transaction.City,
		CreatedAt: transaction.CreatedAt,
	}
}

func FromAccountModelWithTransactionsToResponse(account *models.Account) responses.Account {
	transactionsResponse := make([]*responses.Transaction, 0)
	if account.Transactions != nil {
		for _, transaction := range account.Transactions {
			t := FromTransactionModelToResponse(transaction)
			transactionsResponse = append(transactionsResponse, &t)
		}
	}

	return responses.Account{
		Id:           account.Id,
		CustomerId:   account.CustomerId,
		Kind:         account.Kind,
		Balance:      account.Balance,
		City:         account.City,
		Country:      account.Country,
		Currency:     account.Currency,
		Transactions: transactionsResponse,
		CreatedAt:    account.CreatedAt,
		UpdatedAt:    account.UpdatedAt,
	}
}

func FromReportModelToResponse(reportQtyTX []*models.Report) []responses.Report {
	reportResponse := make([]responses.Report, 0)
	for _, report := range reportQtyTX {
		r := responses.NewReport()
		r.QtyTransactions = report.QtyTransactions
		r.CustomerId = report.CustomerId
		r.CustomerName = report.CustomerName
		r.Month = report.Month
		r.Year = report.Year
		reportResponse = append(reportResponse, r)
	}
	return reportResponse
}

func FromReportBigTransactionsModelToResponse(reportBigTransactions []*models.ReportBigOperation) []responses.ReportBigOperation {
	reportResponse := make([]responses.ReportBigOperation, 0)
	for _, report := range reportBigTransactions {
		r := responses.NewReportBigOperation()
		r.CustomerName = report.CustomerName
		r.CustomerId = report.CustomerId
		r.CreatedAt = report.CreatedAt
		r.City = report.City
		r.OriginalCity = report.OriginalCity
		r.Amount = report.Amount
		reportResponse = append(reportResponse, r)
	}
	return reportResponse
}
