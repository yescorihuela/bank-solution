package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/application/usecases"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/http/api/validators"
	validator "github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/http/api/validators"
)

type TransactionHandlerConfig struct {
	Logger             *logrus.Logger
	Validator          *validator.Validator
	TransactionUseCase usecases.TransactionUseCase
	AccountUseCase     usecases.AccountUseCase
}

type TransactionHandler struct {
	logger             *logrus.Logger
	validator          *validator.Validator
	transactionUseCase usecases.TransactionUseCase
	accountUseCase     usecases.AccountUseCase
}

func NewTransactionHandler(cfg TransactionHandlerConfig) *TransactionHandler {
	return &TransactionHandler{
		logger:             cfg.Logger,
		validator:          cfg.Validator,
		transactionUseCase: cfg.TransactionUseCase,
		accountUseCase:     cfg.AccountUseCase,
	}
}

func (th *TransactionHandler) Deposit(ctx *gin.Context) {

}

func (th *TransactionHandler) WithDraw(ctx *gin.Context) {

}

func (th *TransactionHandler) GetTransactionsByAccountId(ctx *gin.Context) {
	customerId := strings.TrimSpace(ctx.Param("customer_id"))
	accountId := strings.TrimSpace(ctx.Param("account_id"))

	validator := validators.NewValidator()

	validator.Check(customerId != "", "customer_id", "customer_id must be an valid url param")
	validator.Check(accountId != "", "accountId", "accountId must be an valid url param")

	if !validator.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": validator.Errors})
		return
	}

	_, err := th.accountUseCase.GetById(ctx, customerId, accountId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": err})
		return
	}

}

func (th *TransactionHandler) GetTransactionsByAccountIdAndMonth(ctx *gin.Context) {

}
