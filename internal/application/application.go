package application

import (
	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/http/api/handlers"
)

type Application struct {
	accountHandler     *handlers.AccountHandler
	customerHandler    *handlers.CustomerHandler
	transactionHandler *handlers.TransactionHandler
	router             *gin.Engine
	// config             utils.Config
}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) RegisterRoutes() {
	v1 := app.router.Group("/api/v1")
	v1.POST("/customers", app.customerHandler.Create)
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
	err := app.router.Run(":8080") // change with config
	return err
}
