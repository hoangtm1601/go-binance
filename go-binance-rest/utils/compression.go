package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"fmt"
)

// DecodeNatsResponse decodes a gzip-compressed and gob-encoded NATS response
func DecodeNatsResponse[T any](data []byte) (T, error) {
	var result T

	// Decompress the gzip data
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return result, fmt.Errorf("failed to decompress response data: %w", err)
	}
	defer reader.Close()

	// Decode the gob data
	decoder := gob.NewDecoder(reader)
	if err := decoder.Decode(&result); err != nil {
		return result, fmt.Errorf("failed to decode response data: %w", err)
	}

	return result, nil
}
