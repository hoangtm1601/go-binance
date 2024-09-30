package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hoangtm1601/go-binance-rest/internal/api/services"
	"github.com/hoangtm1601/go-binance-rest/internal/middleware"
	"github.com/hoangtm1601/go-binance-rest/internal/models"
	"github.com/hoangtm1601/go-binance-rest/internal/models/dto"
	"net/http"
)

type UserController struct {
	service services.UserServiceInterface
}

func NewUserController(service services.UserServiceInterface) *UserController {
	return &UserController{service}
}

// ListUsers godoc
//
//		@Summary		ListUsers
//		@Description	ListUsers
//		@Tags			users
//		@Accept			json
//		@Produce		json
//		@Param			payload	query		dto.ListUserDto			false	"ListOrders payload"
//		@Param 			_ 		query 		dto.PaginationDto 		false 	"PaginationDto"
//		@Success		200	{object}		dto.UserListResponse
//	 	@Security		Bearer
//		@Router			/users/list [get]
func (uc *UserController) ListUsers(c *gin.Context) {
	var payload dto.ListUserDto
	pagination := c.MustGet(middleware.Pagination).(dto.PaginationDto)
	payload.Email = c.Query("email")
	payload.Name = c.Query("name")
	payload.Role = c.Query("role")
	payload.Provider = c.Query("provider")

	users, total, err := uc.service.ListUsers(payload, pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.ToUserListResponse(users)
	response.Total = total

	c.JSON(http.StatusOK, response)
}

// GetMe godoc
//
//		@Summary		GetMe
//		@Description	GetMe
//		@Tags			users
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	dto.UserResponse
//		@Failure 		500 {string} 	string 				"an error occurred during the modification"
//	 	@Security		Bearer
//		@Router			/users/me [get]
func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet(middleware.CurrentUser).(models.User)
	userResponse := &dto.UserResponse{
		ID:        currentUser.ID,
		Name:      currentUser.Name,
		Email:     currentUser.Email,
		Photo:     currentUser.Photo,
		Role:      currentUser.Role,
		Provider:  currentUser.Provider,
		CreatedAt: currentUser.CreatedAt,
		UpdatedAt: currentUser.UpdatedAt,
	}
	fmt.Printf("%v", userResponse)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}

// GetUser godoc
//
//		@Summary		GetUser
//		@Description	GetUser
//		@Tags			users
//		@Accept			json
//		@Produce		json
//		@Param			id	path		string	true	"User ID"
//		@Success		200	{object}	dto.UserResponse
//		@Failure 		500 {string} 	string 				"an error occurred during the modification"
//	 	@Security		Bearer
//		@Router			/users/{id} [get]
func (uc *UserController) GetUser(ctx *gin.Context) {
	var readUserRequest dto.ReadUserRequest
	if err := ctx.ShouldBindUri(&readUserRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := uc.service.ReadUser(uint(readUserRequest.ID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, dto.ToUserResponse(user))
}
