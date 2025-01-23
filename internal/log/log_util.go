package log_middleware

import (
	"context"
	"log/slog"
)

func LogInfo(response Response, logger *slog.Logger, message string) {
	logger.LogAttrs(context.Background(), slog.LevelInfo, message, slog.String("method", response.Method), slog.String("path", response.Url), slog.Int("status", response.Status), slog.String("messages", response.Message), slog.String("time_taken", response.TimeTaken))
}

func LogError(response Response, logger *slog.Logger, message string) {
	logger.LogAttrs(context.Background(), slog.LevelError, message, slog.String("method", response.Method), slog.String("path", response.Url), slog.Int("status", response.Status), slog.String("messages", response.Message), slog.String("time_taken", response.TimeTaken))
}

func LogWarn(response Response, logger *slog.Logger, message string) {
	logger.LogAttrs(context.Background(), slog.LevelWarn, message, slog.String("method", response.Method), slog.String("path", response.Url), slog.Int("status", response.Status), slog.String("messages", response.Message), slog.String("time_taken", response.TimeTaken))
}
