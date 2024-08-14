package requests

type Account struct {
	Balance  string `json:"balance"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Currency int    `json:"currency"`
}

func NewAccount() *Account {
	return &Account{}
}
