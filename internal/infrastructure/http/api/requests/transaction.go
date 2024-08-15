package requests

type TransactionRequest struct {
	Amount float64 `json:"amount" binding:"required,gt=0"`
	Kind   string  `json:"kind" binding:"numeric,required,gte=0,max=1"`
	City   string  `json:"city" binding:"required"`
}

func NewTransactionRequest() TransactionRequest {
	return TransactionRequest{}
}
