package repositories

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/constants"
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
		SELECT COUNT(t.id) AS qty_transactions,
					c.name AS customer_name,
					c.id AS customer_id,
					EXTRACT(YEAR
									FROM t.created_at::date) AS YEAR,
					EXTRACT(MONTH
									FROM t.created_at::date) AS MONTH
		FROM accounts a
		INNER JOIN transactions t ON(a.id = t.account_id)
		INNER JOIN customers c ON (a.customer_id = c.id)
		WHERE EXTRACT(YEAR
									FROM t.created_at::date) = $1
			AND EXTRACT(MONTH
									FROM t.created_at::date) = $2
		GROUP BY customer_name,
						c.id,
						YEAR,
						MONTH
		ORDER BY qty_transactions DESC
`)

var getBigOperationsCustomers = shared.Compact(`
		SELECT t.city,
					a.city as original_city,
					t.amount,
					c.name AS customer_name,
					a.customer_id,
					t.created_at
		FROM transactions t
		INNER JOIN accounts a ON (t.city != a.city)
		INNER JOIN customers c ON (a.customer_id = c.id)
		WHERE t.amount > $1
			AND extract(YEAR
									FROM t.created_at::date) = $2
			AND extract(MONTH
									FROM t.created_at::date) = $3
		ORDER BY t.created_at DESC,
						t.amount DESC
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
			rrp.logger.Errorf("Failing ReportRepositoryPostgresql.GetTransactionsByCustomers method fetching transactions rows %s", err.Error())
			return reportModels, err
		}
		reportModels = append(reportModels, r)
	}

	return reportModels, nil
}

func (rrp *ReportRepositoryPostgresql) GetBigTransactionsOutSide(ctx context.Context, month, year int) ([]*models.ReportBigOperation, error) {
	rrp.logger.Info("Starting ReportRepositoryPostgresql.GetBigTransactionsOutSide method")
	reportBigOperationModels := make([]*models.ReportBigOperation, 0)
	resultSet, err := rrp.db.Query(ctx, getBigOperationsCustomers, constants.LIMIT_TRANSACTION_REPORT, year, month)
	if err != nil {
		return nil, err
	}

	for resultSet.Next() {
		r := models.NewReportBigOperation()
		err := resultSet.Scan(
			&r.City,
			&r.OriginalCity,
			&r.Amount,
			&r.CustomerName,
			&r.CustomerId,
			&r.CreatedAt,
		)
		if err != nil {
			rrp.logger.Errorf("Failing ReportRepositoryPostgresql.GetBigTransactionsOutSide method fetching transactions rows %s", err.Error())
			return reportBigOperationModels, err
		}
		reportBigOperationModels = append(reportBigOperationModels, r)
	}

	return reportBigOperationModels, nil
}
