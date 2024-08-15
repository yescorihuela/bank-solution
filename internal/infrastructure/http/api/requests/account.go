package requests

type AccountRequest struct {
	Balance  float64 `json:"balance"`
	Kind     int     `json:"kind"`
	City     string  `json:"city"`
	Country  string  `json:"country"`
	Currency int     `json:"currency"`
}

func NewAccountRequest() AccountRequest {
	return AccountRequest{}
}
