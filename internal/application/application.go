package application

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/http/api/handlers"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/shared/utils"
)

type Application struct {
	accountHandler     *handlers.AccountHandler
	customerHandler    *handlers.CustomerHandler
	transactionHandler *handlers.TransactionHandler
	router             *gin.Engine
	config             utils.Config
	logger             *logrus.Logger
}

func NewApplication(
	accountHandler *handlers.AccountHandler,
	customerHandler *handlers.CustomerHandler,
	transactionHandler *handlers.TransactionHandler,
	router *gin.Engine,
	logger *logrus.Logger,
	config utils.Config,
) *Application {
	return &Application{
		accountHandler:     accountHandler,
		customerHandler:    customerHandler,
		transactionHandler: transactionHandler,
		router:             router,
		logger:             logger,
		config:             config,
	}
}

func (app *Application) RegisterRoutes() {
	v1 := app.router.Group("/api/v1")
	v1.POST("/customers", app.customerHandler.Create)
	v1.GET("/customers/:customer_id", app.customerHandler.GetByCustomerId)
	v1.POST("/customers/:customer_id/accounts/", app.accountHandler.Create)
	v1.GET("/customers/:customer_id/accounts/:account_id", app.accountHandler.GetAccountById)
	v1.GET("/customers/:customer_id/accounts/:account_id/latest_transactions", app.accountHandler.GetLastTransactionsByAccountId)
	v1.POST("/customers/:customer_id/accounts/:account_id/deposit", app.transactionHandler.Deposit)
	v1.POST("/customers/:customer_id/accounts/:account_id/withdrawl", app.transactionHandler.WithDraw)
}

func (app *Application) Bootstrapping() {
	app.RegisterRoutes()
}

func (app *Application) Run() error {
	app.Bootstrapping()
	err := app.router.Run(fmt.Sprintf(":%d", app.config.AppHTTPPort))
	return err
}
