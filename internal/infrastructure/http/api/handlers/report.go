package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/application/usecases"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/http/api/validators"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/infrastructure/mappers"
)

type ReportHandlerConfig struct {
	Logger        *logrus.Logger
	ReportUseCase usecases.ReportUseCase
}
type ReportHandler struct {
	logger        *logrus.Logger
	reportUseCase usecases.ReportUseCase
}

func NewReportHandler(cfg ReportHandlerConfig) *ReportHandler {
	return &ReportHandler{
		logger:        cfg.Logger,
		reportUseCase: cfg.ReportUseCase,
	}
}

func (rh *ReportHandler) GetMonthlyTransactionsByCustomers(ctx *gin.Context) {
	var (
		year  int = time.Now().Year()
		month int = int(time.Now().Month())
	)
	validator := validators.NewValidator()

	if strings.TrimSpace(ctx.Query("month")) != "" {
		month, _ = strconv.Atoi(ctx.Query("month"))
		validator.Check(month > 0, "month", "month must be an positive number into query string param")
	}

	if strings.TrimSpace(ctx.Query("year")) != "" {
		year, _ = strconv.Atoi(ctx.Query("year"))
		validator.Check(year > 0, "year", "year must be an positive number into query string param")
	}

	if !validator.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validator.Errors})
		return
	}
	fmt.Println("month, year => ", month, year)
	reportModel, err := rh.reportUseCase.GetTransactionsByCustomers(ctx, month, year)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}
	reportResponse := mappers.FromReportModelToResponse(reportModel)
	ctx.JSON(http.StatusOK, gin.H{"data": reportResponse})
}
