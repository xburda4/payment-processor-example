package main

import (
	"context"
	"fintech-proj/api"
	"fintech-proj/api/config"
	"fintech-proj/util/logger"
	"fmt"
	"log/slog"
	"net/http"
)

const (
	defaultEnvFile = ".env"
)

func main() {
	cfg, err := config.NewConfig(defaultEnvFile)
	if err != nil {
		panic(fmt.Errorf("parsing config: %w", err))
	}

	logger.Initialize(cfg.LogLevel)

	ctx := context.Background()
	addr := fmt.Sprintf(":%d", cfg.Port)

	logger.Default().InfoContext(ctx, "starting application",
		slog.String("addr", addr),
	)

	controller, err := api.NewController(
		"0.0.0",
		"local",
	)
	if err := http.ListenAndServe(addr, controller); err != nil {
		logger.Default().ErrorContext(ctx, "error running http server")
	}
}
