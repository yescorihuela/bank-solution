package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/application/usecases"
	validator "github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/http/api/validators"
)

type AccountHandlerConfig struct {
	logger         *logrus.Logger
	validator      *validator.Validator
	accountUseCase *usecases.AccountUseCase
}

type AccountHandler struct {
	logger         *logrus.Logger
	validator      *validator.Validator
	accountUseCase *usecases.AccountUseCase
}

func NewAccountHandler(cfg AccountHandlerConfig) *AccountHandler {
	return &AccountHandler{
		logger:         cfg.logger,
		validator:      cfg.validator,
		accountUseCase: cfg.accountUseCase,
	}
}

func (ah *AccountHandler) Create(ctx *gin.Context) {}

func (ah *AccountHandler) GetAccountById(ctx *gin.Context) {}

func (ah *AccountHandler) GetLastTransactionsByAccountId(ctx *gin.Context) {}
