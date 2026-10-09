package logger

import (
	"log/slog"
	"os"
)

// New creates and returns a configured structured logger.
func New(env string) *slog.Logger {

	var handler slog.Handler

	switch env {
	case "development":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		})

	case "production":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelInfo,
			AddSource: false,
		})

	default:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		})
	}

	return slog.New(handler)
}
