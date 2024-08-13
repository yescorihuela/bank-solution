package entities

import "time"

const (
	WithDrawal = iota
	Deposit
)

type Transaction struct {
	Id        int64
	AccountId string
	Kind      string
	Status    string
	City      string
	CreatedAt time.Time
}
