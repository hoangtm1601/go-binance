package middleware

import (
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

func LoggingMiddleware(handler nats.MsgHandler) nats.MsgHandler {
	return func(msg *nats.Msg) {
		start := time.Now()
		log.Printf("MicroserviceMiddleware: Received message on subject: %s, payload: %s", msg.Subject, msg.Data)
		handler(msg)
		log.Printf("MicroserviceMiddleware: Finished processing message on subject: %s, took: %v", msg.Subject, time.Since(start))
	}
}
