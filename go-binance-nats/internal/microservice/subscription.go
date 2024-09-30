package microservice

import (
	"github.com/hoangtm1601/go-binance-nats/internal/api/controllers"
	"github.com/hoangtm1601/go-binance-nats/internal/api/repositories"
	"github.com/hoangtm1601/go-binance-nats/internal/api/services"
	"github.com/hoangtm1601/go-binance-nats/internal/initializers"
	"log"
)

func SetupMicroserviceSubscription() {
	// Initialize repositories, services, and controllers
	candleRepo := repositories.NewCandleRepository(initializers.DB)
	candleService := services.NewCandleService(candleRepo)
	candleController := controllers.NewCandleController(candleService, initializers.GetNatsConnection())

	// Setup subscriptions
	err := candleController.SetupSubscriptions()
	if err != nil {
		log.Fatalf("Failed to setup subscriptions: %v", err)
	}
}
