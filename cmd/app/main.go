package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/application"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/application/usecases"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/databases/postgresql"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/http/api/handlers"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/repositories"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/shared/utils"
)

func main() {
	config, err := utils.LoadConfig("../../")
	if err != nil {
		panic(err)
	}
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	db, err := postgresql.NewPostgresDBConnection(config)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	pgAccountRepository := repositories.NewAccountRepositoryPostgresql(db, logger)
	pgCustomerRepository := repositories.NewCustomerRepositoryPostgresql(db, logger)
	pgTransactionRepository := repositories.NewTransactionRepositoryPostgresql(db, logger)
	pgReportRepository := repositories.NewReportRepositoryPostgresql(db, logger)

	accountUseCase := usecases.NewAccountUseCase(
		logger,
		pgAccountRepository,
		pgCustomerRepository,
		pgTransactionRepository,
	)

	customerUseCase := usecases.NewCustomerUseCase(
		logger,
		pgAccountRepository,
		pgCustomerRepository,
		pgTransactionRepository,
	)

	transactionUseCase := usecases.NewTransactionUseCase(
		logger,
		pgAccountRepository,
		pgCustomerRepository,
		pgTransactionRepository,
	)

	reportUseCase := usecases.NewReportUseCase(
		logger,
		pgReportRepository,
	)

	accountHandler := handlers.NewAccountHandler(
		handlers.AccountHandlerConfig{
			Logger:         logger,
			AccountUseCase: accountUseCase,
		},
	)

	customerHandler := handlers.NewCustomerHandler(
		handlers.CustomerHandlerConfig{
			Logger:          logger,
			CustomerUseCase: customerUseCase,
		},
	)

	transactionHandler := handlers.NewTransactionHandler(
		handlers.TransactionHandlerConfig{
			Logger:             logger,
			TransactionUseCase: transactionUseCase,
		},
	)

	reportHandler := handlers.NewReportHandler(
		handlers.ReportHandlerConfig{
			Logger:        logger,
			ReportUseCase: reportUseCase,
		},
	)

	app := application.NewApplication(
		accountHandler,
		customerHandler,
		transactionHandler,
		reportHandler,
		gin.Default(),
		logger,
		config,
	)

	if err := app.Run(); err != nil {
		panic(err)
	}

}
