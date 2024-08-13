package entities

import "time"

type Customer struct {
	Id        string
	Name      string
	Kind      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
