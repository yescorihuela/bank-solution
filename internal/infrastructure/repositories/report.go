package repositories

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/repositories"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/shared"
)

type ReportRepositoryPostgresql struct {
	db     *pgxpool.Pool
	logger *logrus.Logger
}

func NewReportRepositoryPostgresql(
	db *pgxpool.Pool,
	logger *logrus.Logger,
) repositories.ReportRepository {
	return &ReportRepositoryPostgresql{
		db:     db,
		logger: logger,
	}
}

var getCustomerByTransactionsQtyMonthly = shared.Compact(`
	SELECT 
		COUNT(t.id) as qty_transactions,
		c.name as customer_name,
		c.id as customer_id,
		EXTRACT(year FROM t.created_at::date) as year,
		EXTRACT(month from t.created_at::date) as month
	FROM accounts a 
	INNER JOIN 
		transactions t ON(a.id = t.account_id)
	INNER JOIN 
		customers c ON (a.customer_id = c.id)
	WHERE 
		EXTRACT(year from t.created_at::date) = $1 and EXTRACT(month from t.created_at::date) = $2
	GROUP BY customer_name, c.id, year, month
	ORDER BY qty_transactions DESC
`)

func (rrp *ReportRepositoryPostgresql) GetTransactionsByCustomers(ctx context.Context, month, year int) ([]*models.Report, error) {
	rrp.logger.Info("Starting ReportRepositoryPostgresql.GetTransactionsByCustomers method")
	reportModels := make([]*models.Report, 0)
	resultSet, err := rrp.db.Query(ctx, getCustomerByTransactionsQtyMonthly, year, month)
	if err != nil {
		rrp.logger.Errorf("Failing ReportRepositoryPostgresql.GetTransactionsByCustomers method querying transactions %s", err)
		return nil, err
	}

	for resultSet.Next() {
		r := models.NewReportModel()
		err := resultSet.Scan(
			&r.QtyTransactions,
			&r.CustomerName,
			&r.CustomerId,
			&r.Year,
			&r.Month,
		)
		if err != nil {
			rrp.logger.Errorf("Failing ReportRepositoryPostgresql.GetTransactionsByCustomers method fetching transactions rows %s", err)
			return reportModels, err
		}
		reportModels = append(reportModels, r)
	}

	return reportModels, nil
}
