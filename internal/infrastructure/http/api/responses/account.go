package responses

import "github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"

type Account struct{}

func NewAccount(account *models.Account) *Account {
	return &Account{}
}
