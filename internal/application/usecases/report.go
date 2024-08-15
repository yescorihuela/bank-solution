package usecases

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/repositories"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/models"
)

type ReportUseCase interface {
	GetTransactionsByCustomers(ctx context.Context, month, year int) ([]*models.Report, error)
}

type reportUseCase struct {
	logger           *logrus.Logger
	reportRepository repositories.ReportRepository
}

func NewReportUseCase(
	logger *logrus.Logger,
	reportRepository repositories.ReportRepository,
) ReportUseCase {
	return &reportUseCase{
		logger:           logger,
		reportRepository: reportRepository,
	}
}

func (ruc reportUseCase) GetTransactionsByCustomers(ctx context.Context, month, year int) ([]*models.Report, error) {
	ruc.logger.Info("Starting reportUseCase.GetTransactionsByCustomers method")
	reportModels, err := ruc.reportRepository.GetTransactionsByCustomers(ctx, month, year)
	if err != nil {
		ruc.logger.Error("Error during access to reportUseCase in reportRepository.GetTransactionsByCustomers method")
		return nil, err
	}
	ruc.logger.Info("reportUseCase.GetTransactionsByCustomers executed successfully")
	return reportModels, nil
}
