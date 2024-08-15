package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/application/usecases"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/constants"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/http/api/requests"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/http/api/validators"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/mappers"
)

type AccountHandlerConfig struct {
	Logger         *logrus.Logger
	AccountUseCase usecases.AccountUseCase
}

type AccountHandler struct {
	logger         *logrus.Logger
	accountUseCase usecases.AccountUseCase
}

func NewAccountHandler(cfg AccountHandlerConfig) *AccountHandler {
	return &AccountHandler{
		logger:         cfg.Logger,
		accountUseCase: cfg.AccountUseCase,
	}
}

func (ah *AccountHandler) Create(ctx *gin.Context) {
	req := requests.NewAccountRequest()
	customerId := ctx.Param("customer_id")
	validator := validators.NewValidator()
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validator.Check(req.Balance >= 0, "balance", "balance field must be equal or greater than zero")
	validator.Check(req.City != "", "city", "city field is required")
	validator.Check(req.Country != "", "country", "country field is required")
	validator.Check(validators.AllowedValue(req.Currency, constants.Currencies...), "currency", fmt.Sprintf("currency field allows the following currencies %v", constants.Currencies))
	if !validator.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validator.Errors})
		return
	}

	accountEntity := mappers.FromAccountRequestToEntity(req, customerId)

	accountModel, err := ah.accountUseCase.Insert(ctx, accountEntity)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accountResponse := mappers.FromAccountModelToResponse(accountModel)
	ctx.JSON(http.StatusCreated, accountResponse)
}

func (ah *AccountHandler) GetAccountById(ctx *gin.Context) {
	customerId := ctx.Param("customer_id")
	accountId := ctx.Param("account_id")
	validator := validators.NewValidator()

	validator.Check(customerId != "", "customer_id", "customer_id must be an valid url param")
	validator.Check(accountId != "", "account_id", "account_id must be an valid url param")

	if !validator.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validator.Errors})
		return
	}

	accountModel, err := ah.accountUseCase.GetById(ctx, customerId, accountId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	accountResponse := mappers.FromAccountModelToResponse(accountModel)
	ctx.JSON(http.StatusOK, accountResponse)
}

func (ah *AccountHandler) GetLastTransactionsByAccountId(ctx *gin.Context) {
	customerId := ctx.Param("customer_id")
	accountId := ctx.Param("account_id")
	validator := validators.NewValidator()
	var lastTransactionsNumber int = constants.LAST_TRANSACTIONS_NUMBER_BY_DEFAULT
	if strings.TrimSpace(ctx.Query("qty_tx")) != "" {
		n, _ := strconv.Atoi(ctx.Query("qty_tx"))	
		validator.Check(n > 0, "qty_tx", "customer_id must be an positive number into query string param")
		lastTransactionsNumber = n
	}

	validator.Check(customerId != "", "customer_id", "customer_id must be an valid url param")
	validator.Check(accountId != "", "account_id", "account_id must be an valid url param")

	if !validator.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validator.Errors})
		return
	}

	accountModel, err := ah.accountUseCase.GetLastTransactionsById(ctx, lastTransactionsNumber, customerId, accountId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": err})
		return
	}

	accountResponse := mappers.FromAccountModelWithTransactionsToResponse(accountModel)
	ctx.JSON(http.StatusOK, accountResponse)
}

func (ah *AccountHandler) GetAccountWithTransactionsByAccountIdAndMonth(ctx *gin.Context) {
	customerId := ctx.Param("customer_id")
	accountId := ctx.Param("account_id")

	validator := validators.NewValidator()

	var (
		month int
		year  int
	)
	if strings.TrimSpace(ctx.Query("month")) != "" {
		month, _ = strconv.Atoi(ctx.Query("month"))
		validator.Check(month > 0, "month", "month must be an positive number into query string param")
	}

	if strings.TrimSpace(ctx.Query("year")) != "" {
		year, _ = strconv.Atoi(ctx.Query("year"))
		validator.Check(year > 0, "year", "year must be an positive number into query string param")
	}

	validator.Check(customerId != "", "customer_id", "customer_id must be an valid url param")
	validator.Check(accountId != "", "account_id", "account_id must be an valid url param")

	if !validator.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validator.Errors})
		return
	}

	accountModel, err := ah.accountUseCase.GetLastTransactionsByAccountIdAndMonth(ctx, month, year, customerId, accountId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": err})
		return
	}

	accountResponse := mappers.FromAccountModelWithTransactionsToResponse(accountModel)
	ctx.JSON(http.StatusOK, accountResponse)
}
