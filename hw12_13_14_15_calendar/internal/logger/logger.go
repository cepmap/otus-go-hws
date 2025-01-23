package logger

import (
	"bytes"
	"log/slog"
	"strings"
)

func New(level string) *slog.Logger {
	var logLevel = &slog.LevelVar{}

	var logBuffer = new(bytes.Buffer)

	opts := &slog.HandlerOptions{
		Level: logLevel,
	}

	handler := slog.NewJSONHandler(logBuffer, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	level = strings.ToLower(level)

	switch level {
	case "debug":
		logLevel.Set(slog.LevelDebug)
	case "info":
		logLevel.Set(slog.LevelInfo)
	case "warn":
		logLevel.Set(slog.LevelWarn)
	case "error":
		logLevel.Set(slog.LevelError)
	default:
		logLevel.Set(slog.LevelInfo)
	}

	return logger
}
