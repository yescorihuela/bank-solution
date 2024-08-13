package entities

import "time"

type Transaction struct {
	Id        int64
	AccountId string
	Kind      string
	Status    string
	City      string
	CreatedAt time.Time
}
