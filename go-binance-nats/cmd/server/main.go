package main

import (
	"github.com/hoangtm1601/go-binance-nats/internal/microservice"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/context"

	"github.com/hoangtm1601/go-binance-nats/internal/initializers"
)

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	// Initialize your database connection
	initializers.ConnectDB(&config)

	// Initialize Redis
	initializers.InitRedis(&config)

	// Initialize NATs
	natsClient := initializers.ConnectNATS(&config)

	microservice.SetupMicroserviceSubscription()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)
	<-quit
	log.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Perform any cleanup operations here
	if err := natsClient.Drain(); err != nil {
		log.Printf("Error draining NATS connection: %v", err)
	}

	<-ctx.Done()
	log.Println("Server exited")
}
