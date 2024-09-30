package dto

import (
	"time"

	"github.com/hoangtm1601/go-binance-rest/internal/models"
)

type CreateTransactionDTO struct {
	Amount   float64 `form:"amount" json:"amount" binding:"required"`
	Currency string  `form:"currency" json:"currency" binding:"required"`
}

type TransactionResponseDTO struct {
	ID              uint                     `json:"id"`
	UserId          uint                     `json:"user_id"`
	Amount          float64                  `json:"amount"`
	TransactionDate time.Time                `json:"transaction_date"`
	Status          models.TransactionStatus `json:"status"`
	Currency        string                   `json:"currency"`
	Type            models.TransactionType   `json:"type"`
}

type PaginationMetadataDTO struct {
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
	Page    int `json:"page"`
}

type IndexTransactionsResponseDTO struct {
	Transactions []*models.Transaction `json:"transactions"`
	Pagination   PaginationMetadataDTO `json:"pagination"`
}
