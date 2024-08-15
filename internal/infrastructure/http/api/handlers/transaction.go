package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/application/usecases"
	validator "github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/http/api/validators"
)

type TransactionHandlerConfig struct {
	Logger             *logrus.Logger
	Validator          *validator.Validator
	TransactionUseCase usecases.TransactionUseCase
}

type TransactionHandler struct {
	logger             *logrus.Logger
	validator          *validator.Validator
	transactionUseCase usecases.TransactionUseCase
}

func NewTransactionHandler(cfg TransactionHandlerConfig) *TransactionHandler {
	return &TransactionHandler{
		logger:             cfg.Logger,
		validator:          cfg.Validator,
		transactionUseCase: cfg.TransactionUseCase,
	}
}

func (th *TransactionHandler) Deposit(ctx *gin.Context) {}

func (th *TransactionHandler) WithDraw(ctx *gin.Context) {}
