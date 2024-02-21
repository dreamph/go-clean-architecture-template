package main

import (
	applogger "backend/internal/core/logger"
	"backend/internal/core/logger/zap"
	"backend/internal/core/utils"
	"context"
)

func main() {
	logger := zap.NewLogger(&applogger.Options{
		FilePath: "./logxxx.log",
		Level:    "debug",
		Format:   "json",
		ProdMode: false,
		MaxAge:   30,
		MaxSize:  10,
	})
	defer logger.Sync()

	d := map[string]string{
		"requestId": utils.NewID(),
	}

	ctx := applogger.WithValue(context.Background(), d)

	for i := 1; i <= 10000000000000; i++ {
		do(ctx, logger)
	}
}

func do(ctx context.Context, logger applogger.Logger) {
	log := logger.WithLogger(ctx)

	log.Info("hellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohello")
}
