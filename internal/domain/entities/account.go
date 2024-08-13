package entities

import "time"

const (
	USD = iota
	EUR
	CLP
	COP
	MXN
	ARS
	CAD
)

type Account struct {
	Id         string
	CustomerId string
	Balance    float64
	City       string
	Country    string
	Currency   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
