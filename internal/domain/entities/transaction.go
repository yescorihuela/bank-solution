package entities

import (
	"time"
)

type Transaction struct {
	Id        string
	AccountId string
	Balance   float64
	Kind      int
	Status    int
	City      string
	CreatedAt time.Time
}
