package utils

import (
	"log/slog"
	"os"

	"youpiteron.dev/wildlands-backend/internal/api"
)

type ConsoleLogger struct {
	logger *slog.Logger
}

var _ api.Logger = (*ConsoleLogger)(nil)

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{logger: slog.New(slog.NewJSONHandler(os.Stdout, nil))}
}

func (l *ConsoleLogger) Info(message string, args ...any) {
	l.logger.Info(message, args...)
}

func (l *ConsoleLogger) Warn(message string, args ...any) {
	l.logger.Warn(message, args...)
}

func (l *ConsoleLogger) Error(message string, args ...any) {
	l.logger.Error(message, args...)
}
