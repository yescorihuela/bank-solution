package handlers

import (
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

type TransactionHandlerConfig struct {
	Logger             *logrus.Logger
	Validator          *validators.Validator
	TransactionUseCase usecases.TransactionUseCase
	AccountUseCase     usecases.AccountUseCase
}

type TransactionHandler struct {
	logger             *logrus.Logger
	validator          *validators.Validator
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

func (th *TransactionHandler) CreateTransaction(ctx *gin.Context) {
	req := requests.NewTransactionRequest()
	customerId := strings.TrimSpace(ctx.Param("customer_id"))
	accountId := strings.TrimSpace(ctx.Param("account_id"))

	validator := validators.NewValidator()

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	kind, _ := strconv.Atoi(req.Kind)

	validator.Check(customerId != "", "customer_id", "customer_id must be an valid url param")
	validator.Check(accountId != "", "accountId", "accountId must be an valid url param")
	validator.Check(validators.AllowedValue(kind, []int{constants.Deposit, constants.WithDrawal}...), "kind", "the field kind must be 0 for withdrawal or 1 for deposit")

	if !validator.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": validator.Errors})
		return
	}

	transactionEntity := mappers.FromTransactionRequestToEntity(req, accountId)

	transactionModel, err := th.transactionUseCase.CreateTransaction(ctx, *transactionEntity)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	transactionResponse := mappers.FromTransactionModelToResponse(transactionModel)
	ctx.JSON(http.StatusCreated, gin.H{"data": transactionResponse})
}
