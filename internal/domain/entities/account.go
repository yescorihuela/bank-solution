package entities

import (
	"time"
)

type Account struct {
	Id         string
	Kind       int
	CustomerId string
	Balance    float64
	City       string
	Country    string
	Currency   int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
