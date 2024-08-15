package responses

import "time"

type Transaction struct {
	Id        string    `json:"id"`
	AccountId string    `json:"account_id"`
	Amount    float64   `json:"amount"`
	Kind      int       `json:"kind"`
	Status    int       `json:"status"`
	City      string    `json:"city"`
	CreatedAt time.Time `json:"created_at"`
}

func NewTransaction() *Transaction {
	return &Transaction{}
}
