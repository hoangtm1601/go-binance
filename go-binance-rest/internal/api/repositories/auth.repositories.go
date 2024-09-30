package repositories

import (
	"github.com/hoangtm1601/go-binance-rest/internal/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func (r *AuthRepository) GetDB() *gorm.DB {
	return r.db
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *AuthRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *AuthRepository) CreateWithTx(tx *gorm.DB, user *models.User) error {
	return tx.Session(&gorm.Session{FullSaveAssociations: true}).Create(user).Error
}
