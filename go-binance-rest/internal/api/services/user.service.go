package services

import (
	"github.com/hoangtm1601/go-binance-rest/internal/api/repositories"
	"github.com/hoangtm1601/go-binance-rest/internal/models"
	"github.com/hoangtm1601/go-binance-rest/internal/models/dto"
)

type UserServiceInterface interface {
	ListUsers(dto dto.ListUserDto, pagination dto.PaginationDto) ([]models.User, int, error)
	ReadUser(id uint) (*models.User, error)
}

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) ListUsers(dto dto.ListUserDto, pagination dto.PaginationDto) ([]models.User, int, error) {
	return s.repo.List(dto, pagination)
}

func (s *UserService) ReadUser(id uint) (*models.User, error) {
	user, err := s.repo.Read(id)

	return user, err
}
