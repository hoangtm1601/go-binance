package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangtm1601/go-binance-rest/internal/api/services"
	"github.com/hoangtm1601/go-binance-rest/internal/middleware"
	"github.com/hoangtm1601/go-binance-rest/internal/models"
	"github.com/hoangtm1601/go-binance-rest/internal/models/dto"
)

type PaymentController struct {
	service *services.PaymentService
}

func NewPaymentController(service *services.PaymentService) *PaymentController {
	return &PaymentController{service: service}
}

// CreatePayment godoc
// @Summary Create a new payment
// @Description Create a new payment transaction for the current user
// @Tags payments
// @Accept json
// @Produce json
// @Param payload body dto.CreateTransactionDTO true "CreatePayment payload"
// @Success 200 {object} dto.TransactionResponseDTO
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security Bearer
// @Router /payments [post]
func (pc *PaymentController) CreatePayment(ctx *gin.Context) {
	var payload *dto.CreateTransactionDTO

	if err := ctx.ShouldBind(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	currentUser := ctx.MustGet(middleware.CurrentUser).(models.User)

	tx, err := pc.service.CreatePayment(currentUser.ID, payload)

	if tx != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"transaction": tx}})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

// IndexPayment godoc
// @Summary List user payments
// @Description Retrieve a paginated list of payment transactions for the current user
// @Tags payments
// @Accept json
// @Produce json
// @Param pagination query dto.PaginationDto false "Pagination parameters"
// @Success 200 {object} dto.IndexTransactionsResponseDTO
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security Bearer
// @Router /payments [get]
func (pc *PaymentController) IndexPayment(ctx *gin.Context) {
	pagination := ctx.MustGet(middleware.Pagination).(*dto.PaginationDto)
	currentUser := ctx.MustGet(middleware.CurrentUser).(models.User)
	transactions, total, err := pc.service.UserIndexPayment(currentUser.ID, pagination)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	response := dto.IndexTransactionsResponseDTO{
		Transactions: transactions,
		Pagination: dto.PaginationMetadataDTO{
			Page:    pagination.Page,
			PerPage: pagination.PerPage,
			Total:   int(total),
		},
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": response})
}
