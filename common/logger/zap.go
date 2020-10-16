package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "123.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
