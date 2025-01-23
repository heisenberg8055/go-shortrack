package log_middleware

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

type Response struct {
	Method    string
	Url       string
	Status    int
	Message   string
	TimeTaken string
}

func NewLogger() *slog.Logger {
	logger := slog.New(tint.NewHandler(os.Stderr, nil))
	return logger
}
