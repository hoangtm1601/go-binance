package repositories

import (
	"github.com/hoangtm1601/go-binance-rest/internal/models"
	"github.com/hoangtm1601/go-binance-rest/internal/models/dto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) GetDB() *gorm.DB {
	return r.db
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) UpdateUserRole(id uint, role string) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).Update("role", role).Error
}

func (r *UserRepository) CreateWithTx(tx *gorm.DB, user *models.User) error {
	return tx.Session(&gorm.Session{FullSaveAssociations: true}).Create(user).Error
}

func (r *UserRepository) List(dto dto.ListUserDto, pagination dto.PaginationDto) ([]models.User, int, error) {

	var users []models.User
	var total int64

	offset := (pagination.Page - 1) * pagination.PerPage

	clauses := make([]clause.Expression, 0)
	if dto.Email != "" {
		clauses = append(clauses, clause.Like{Column: "email", Value: dto.Email})
	}
	if dto.Name != "" {
		clauses = append(clauses, clause.Like{Column: "name", Value: "%" + dto.Name + "%"})
	}
	if dto.Provider != "" {
		clauses = append(clauses, clause.Eq{Column: "provider", Value: dto.Provider})
	}
	if dto.Role != "" {
		clauses = append(clauses, clause.Eq{Column: "role", Value: dto.Provider})
	}

	err := r.db.Model(&models.User{}).Clauses(clauses...).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.
		Model(&models.User{}).
		Clauses(clauses...).
		Offset(offset).
		Limit(pagination.PerPage).
		Find(&users).
		Error
	if err != nil {
		return nil, 0, err
	}

	return users, int(total), nil
}

func (r *UserRepository) Read(id uint) (*models.User, error) {
	var user *models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
