package responses

import "github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"

type Transaction struct{}

func NewTransaction(customer *models.Transaction) *Transaction {
	return &Transaction{}
}
