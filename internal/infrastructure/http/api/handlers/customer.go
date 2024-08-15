package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/application/usecases"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/domain/constants"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/http/api/requests"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/http/api/validators"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/mappers"
)

type CustomerHandlerConfig struct {
	Logger          *logrus.Logger
	CustomerUseCase usecases.CustomerUseCase
}
type CustomerHandler struct {
	logger          *logrus.Logger
	customerUseCase usecases.CustomerUseCase
}

func NewCustomerHandler(cfg CustomerHandlerConfig) *CustomerHandler {
	return &CustomerHandler{
		logger:          cfg.Logger,
		customerUseCase: cfg.CustomerUseCase,
	}
}

func (ch *CustomerHandler) Create(ctx *gin.Context) {
	req := requests.NewCustomerRequest()
	validator := validators.NewValidator()
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	validator.Check(req.Name != "", "name", "name field is required")
	validator.Check(validators.AllowedValue[int](req.Kind, []int{constants.Individual, constants.Organization}...), "kind", "the kind of customer must be either Individual(0) or Organization(1)")
	if !validator.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": validator.Errors})
		return
	}

	customerEntity := mappers.FromCustomerRequestToEntity(req)

	customerModel, err := ch.customerUseCase.Insert(ctx, customerEntity)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	customerResponse := mappers.FromCustomerModelToResponse(customerModel)

	ctx.JSON(http.StatusCreated, customerResponse)
}

func (ch *CustomerHandler) GetByCustomerId(ctx *gin.Context) {
	customerId := ctx.Param("customer_id")
	validator := validators.NewValidator()
	validator.Check(customerId != "", "customer_id", "customer_id must be an valid url param")
	if !validator.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": validator.Errors})
		return
	}
	customerModel, err := ch.customerUseCase.GetById(ctx, customerId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		return
	}

	customerResponse := mappers.FromCustomerModelToResponse(customerModel)

	ctx.JSON(http.StatusOK, customerResponse)
}
