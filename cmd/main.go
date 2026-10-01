package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ppMTUCI/harvester/internal/config"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		slog.Error(fmt.Sprintf("Cannot init config: %s", err))
	}

	logger := slog.New(slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: cfg.PrepareLogLevel(),
		},
	))

	logger.Info(fmt.Sprintf("App started with config:%+v", cfg))
}
