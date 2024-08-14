package requests

type Transaction struct {
	Kind   int    `json:"kind"`
	Status int    `json:"status"`
	City   string `json:"city"`
}

func NewTransaction() *Transaction {
	return &Transaction{}
}
