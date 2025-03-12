package main

import (
	"log/slog"
	"os"

	"example.com/url-shorterner/internal/config"
)

func main() {
	// init config - cleanenv
	cfg := config.MustLoad()

	// init logger - slog

	log := setupLogger(cfg.Env)
	log.Info("test logger", slog.String("env", cfg.Env))
	// init storage - sqlite
	// init router - chi
	// run server
}

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))

	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))

	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}

	return log
}
