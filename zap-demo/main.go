package main

import (
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}
	
	logger.Info("Hello from zap logger")
}

/*
{"level":"info","timestamp":"2026-10-21T15:01:13+08:00","caller":"zap-demo/main.go:21","msg":"Hello from zap logger"}
*/