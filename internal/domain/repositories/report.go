package repositories

import (
	"context"

	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"
)

type ReportRepository interface {
	GetTransactionsByCustomers(ctx context.Context, month, year int) ([]*models.Report, error)
}
