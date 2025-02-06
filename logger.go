package gogen

import (
	"os"

	"log/slog"
)

func GetTextLogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	return logger
}

func GetJSONLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	return logger
}
