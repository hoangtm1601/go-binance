package initializers

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
)

var (
	nc   *nats.Conn
	once sync.Once
)

func ConnectNATS(config *Config) *nats.Conn {
	var err error
	once.Do(func() {
		opts := []nats.Option{
			nats.Name("YourAppName"),
			nats.ReconnectWait(time.Second),
			nats.MaxReconnects(-1),
			nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
				log.Printf("Disconnected from NATS: %v", err)
			}),
			nats.ReconnectHandler(func(nc *nats.Conn) {
				log.Printf("Reconnected to NATS: %v", nc.ConnectedUrl())
			}),
			nats.ClosedHandler(func(nc *nats.Conn) {
				log.Println("NATS connection closed")
			}),
		}

		url := fmt.Sprintf("%s:%s", config.NATSHost, config.NATSPort)
		nc, err = nats.Connect(url, opts...)
		if err != nil {
			log.Fatalf("Failed to connect to NATS: %v", err)
			return
		}

		log.Println("Connected to NATS")
	})

	if err != nil {
		log.Fatalf("Error when connecting to NATS: %v", err)
	}

	if nc == nil || !nc.IsConnected() {
		log.Fatalf("NATS connection is not established")
	}

	return nc
}

func GetNatsConnection() *nats.Conn {
	return nc
}
