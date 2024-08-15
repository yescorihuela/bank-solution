package responses

import (
	"time"
)

type Customer struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Kind      int       `json:"kind"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCustomer() *Customer {
	return &Customer{}
}
