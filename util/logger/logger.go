package logger

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func Initialize(level slog.Level) {
	jsonHandler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	})
	logger = slog.New(jsonHandler)
}

func Default() *slog.Logger {
	return logger
}
