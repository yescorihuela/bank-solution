package responses

import (
	"time"

	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"
)

type Account struct {
	Id           string         `json:"id"`
	CustomerId   string         `json:"customer_id,omitempty"`
	Kind         int            `json:"kind"`
	Balance      float64        `json:"balance"`
	City         string         `json:"city"`
	Country      string         `json:"country"`
	Currency     int            `json:"currency"`
	Transactions []*Transaction `json:"transactions"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

func NewAccount(account *models.Account) *Account {
	return &Account{}
}
