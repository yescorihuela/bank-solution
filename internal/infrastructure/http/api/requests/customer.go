package requests

type CustomerRequest struct {
	Name string `json:"name"`
	Kind int    `json:"kind"`
}

func NewCustomerRequest() *CustomerRequest {
	return &CustomerRequest{}
}
