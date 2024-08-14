package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/application/usecases"
	validator "github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/http/api/validators"
)

type CustomerHandlerConfig struct {
	logger          *logrus.Logger
	validator       *validator.Validator
	customerUseCase *usecases.CustomerUseCase
}
type CustomerHandler struct {
	logger          *logrus.Logger
	validator       *validator.Validator
	customerUseCase *usecases.CustomerUseCase
}

func NewCustomerHandler(cfg CustomerHandlerConfig) *CustomerHandler {
	return &CustomerHandler{
		logger:          cfg.logger,
		validator:       cfg.validator,
		customerUseCase: cfg.customerUseCase,
	}
}

func (ch *CustomerHandler) Create(ctx *gin.Context) {}

func (ch *CustomerHandler) GetByCustomerId(ctx *gin.Context) {}
