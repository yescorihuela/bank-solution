package models

import "time"

type Transaction struct {
	Id        int64     `json:"id"`
	AccountId string    `json:"account_id"`
	Kind      string    `json:"kind"`
	Status    string    `json:"status"`
	City      string    `json:"city"`
	CreatedAt time.Time `json:"created_at"`
}
