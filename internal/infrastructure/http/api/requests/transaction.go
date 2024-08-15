package requests

type TransactionRequest struct {
	Amount float64 `json:"amount"`
	Kind   int     `json:"kind"`
	City   string  `json:"city"`
}

func NewTransaction() *TransactionRequest {
	return &TransactionRequest{}
}
