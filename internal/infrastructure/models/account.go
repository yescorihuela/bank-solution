package models

import "time"

type Account struct {
	Id           string         `json:"id"`
	CustomerId   string         `json:"customer_id"`
	Kind         int            `json:"kind"`
	Balance      float64        `json:"balance"`
	City         string         `json:"citiy"`
	Country      string         `json:"country"`
	Currency     int            `json:"currency"`
	Transactions []*Transaction `json:"transactions,omitempty"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

func NewAccountModel() Account {
	return Account{}
}
