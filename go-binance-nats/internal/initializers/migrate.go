package initializers

import (
	"github.com/hoangtm1601/go-binance-nats/internal/models"
)

func Migrate() error {
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	return DB.AutoMigrate(&models.Candle{})
}
