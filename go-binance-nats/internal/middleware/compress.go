package middleware

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"log"

	"github.com/nats-io/nats.go"
)

func CompressAndEncodeMiddleware(handler func(*nats.Msg) (interface{}, error)) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		// Call the original handler
		response, err := handler(msg)
		if err != nil {
			log.Printf("Error in handler: %v", err)
			return
		}

		// Compress and encode the response
		var compressedBuffer bytes.Buffer
		gzipWriter := gzip.NewWriter(&compressedBuffer)
		encoder := gob.NewEncoder(gzipWriter)

		if err := encoder.Encode(response); err != nil {
			log.Printf("Error encoding response: %v", err)
			return
		}

		if err := gzipWriter.Close(); err != nil {
			log.Printf("Error closing gzip writer: %v", err)
			return
		}

		// Send the compressed response back through NATS
		if err := msg.Respond(compressedBuffer.Bytes()); err != nil {
			log.Printf("Error responding with compressed data: %v", err)
		}
	}
}
