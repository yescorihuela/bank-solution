package entities

import (
	"time"
)

type Customer struct {
	Id        string
	Name      string
	Kind      int
	CreatedAt time.Time
	UpdatedAt time.Time
}
