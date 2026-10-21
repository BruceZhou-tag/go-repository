package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	url := "www.google.com"
	logger.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second))
}

/*
{"level":"info","ts":1792565665.8771222,"caller":"zap-demo/main.go:13","msg":"failed to fetch URL","url":"www.google.com","attempt":3,"backoff":1}
*/