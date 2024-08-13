package entities

import (
	"time"

	"github.com/oklog/ulid/v2"
)

const (
	WithDrawal = iota
	Deposit
)

func NewUlid() string {
	return ulid.Make().String()
}

type Transaction struct {
	Id        int64
	AccountId string
	Kind      string
	Status    string
	City      string
	CreatedAt time.Time
}
