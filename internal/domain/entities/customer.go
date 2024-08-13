package entities

import "time"

const (
	Individual = iota
	Organization
)

type Customer struct {
	Id        string
	Name      string
	Kind      string // enum
	CreatedAt time.Time
	UpdatedAt time.Time
}
