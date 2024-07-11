package service

import (
	"log/slog"
	"os"
)

const (
	DEV  = "develop"
	PROD = "production"
)

func SetupLogger(currentEnv string) *slog.Logger {
	var logger *slog.Logger

	switch currentEnv {
	case DEV:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case PROD:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	logger.Info("Starting MOA application", slog.String("env", currentEnv))

	return logger
}
