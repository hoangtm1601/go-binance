package routes

import (
	"github.com/hoangtm1601/go-binance-rest/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hoangtm1601/go-binance-rest/internal/api/controllers"
	"github.com/hoangtm1601/go-binance-rest/internal/api/repositories"
	"github.com/hoangtm1601/go-binance-rest/internal/api/services"
	"github.com/hoangtm1601/go-binance-rest/internal/middleware"
	"gorm.io/gorm"
)

func SetupRoutes(server *gin.Engine, db *gorm.DB) {
	router := server.Group("/api")
	router.GET("/healthcheck", func(ctx *gin.Context) {
		message := "Ping Pong"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthController(authService)

	authRouters := router.Group("/auth")
	{
		authRouters.POST("/register", authController.SignUpUser)
		authRouters.POST("/login", authController.SignInUser)
		authRouters.GET("/refresh", authController.RefreshAccessToken)
	}

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	userRoutes := router.Group("users")
	{
		userRoutes.GET("/list", middleware.BindPagination(), middleware.AuthMiddleware(), middleware.RequireRole(middleware.ADMIN), userController.ListUsers)
		userRoutes.GET("/me", middleware.AuthMiddleware(), userController.GetMe)
		userRoutes.GET("/:id", middleware.AuthMiddleware(), middleware.RequireRole(middleware.ADMIN), userController.GetUser)
	}

	// Add the new Candle routes
	candleService := services.NewCandleService()
	candleController := controllers.NewCandleController(candleService)

	candleRoutes := router.Group("candles").Use(middleware.AuthMiddleware())
	{
		candleRoutes.GET("/indicators", middleware.CacheMiddleware(30*time.Second), middleware.RequireRole(string(models.PAID)), candleController.GetCandlesWithIndicators)
	}

	transactionRepo := repositories.NewTransactionRepository(db)
	paymentService := services.NewPaymentService(transactionRepo, userRepo)
	paymentController := controllers.NewPaymentController(paymentService)

	paymentRoutes := router.Group("payments").Use(middleware.AuthMiddleware())
	{
		paymentRoutes.POST("/", paymentController.CreatePayment)
		paymentRoutes.GET("/", middleware.BindPagination(), paymentController.IndexPayment)
	}
}
