package responses

import "github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"

type Customer struct{}

func NewCustomer(customer *models.Customer) *Customer {
	return &Customer{}
}
