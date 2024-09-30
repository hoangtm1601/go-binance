package services

import (
	"time"

	"github.com/hoangtm1601/go-binance-rest/internal/api/repositories"
	"github.com/hoangtm1601/go-binance-rest/internal/initializers"
	"github.com/hoangtm1601/go-binance-rest/internal/models"
	"github.com/hoangtm1601/go-binance-rest/internal/models/dto"
	"gorm.io/gorm"
)

type PaymentService struct {
	transactionRepo *repositories.TransactionRepository
	userRepo        *repositories.UserRepository
}

func NewPaymentService(transactionRepo *repositories.TransactionRepository, userRepo *repositories.UserRepository) *PaymentService {
	return &PaymentService{transactionRepo: transactionRepo, userRepo: userRepo}
}

func (s *PaymentService) UserIndexPayment(userId uint, pagination *dto.PaginationDto) ([]*models.Transaction, int64, error) {
	var transactions []*models.Transaction
	var total int64

	offset := (pagination.Page - 1) * pagination.PerPage

	countErr := s.transactionRepo.GetDB().Model(&models.Transaction{}).Where("user_id = ?", userId).Count(&total).Error
	if countErr != nil {
		return nil, 0, countErr
	}

	paginateErr := s.transactionRepo.GetDB().Model(&models.Transaction{}).Where("user_id = ?", userId).Limit(pagination.PerPage).Offset(offset).Find(&transactions).Error

	if paginateErr != nil {
		return nil, total, paginateErr
	}

	return transactions, total, nil
}

func (s *PaymentService) CreatePayment(userId uint, payload *dto.CreateTransactionDTO) (*dto.TransactionResponseDTO, error) {
	var createdTransaction models.Transaction

	err := initializers.DB.Transaction(func(tx *gorm.DB) error {
		// Create the transaction and store it in createdTransaction
		createTxErr := tx.Create(&models.Transaction{
			UserId:          userId,
			Amount:          payload.Amount,
			TransactionDate: time.Now(),
			Status:          models.Succeeded,
			Currency:        payload.Currency,
			Type:            models.Charge,
		}).Scan(&createdTransaction).Error

		if createTxErr != nil {
			return createTxErr
		}

		updateErr := tx.Model(&models.User{}).Where("id = ?", userId).Update("role", models.PAID).Error

		if updateErr != nil {
			return updateErr
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Convert the created transaction to DTO and return
	return &dto.TransactionResponseDTO{
		ID:              createdTransaction.ID,
		UserId:          createdTransaction.UserId,
		Amount:          createdTransaction.Amount,
		TransactionDate: createdTransaction.TransactionDate,
		Status:          createdTransaction.Status,
		Currency:        createdTransaction.Currency,
		Type:            createdTransaction.Type,
	}, nil
}
