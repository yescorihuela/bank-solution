package models

import "time"

type Account struct {
	Id         string    `json:"id"`
	CustomerId string    `json:"customer_id"`
	Balance    float64   `json:"balance"`
	City       string    `json:"citiy"`
	Country    string    `json:"country"`
	Currency   string    `json:"currency"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
