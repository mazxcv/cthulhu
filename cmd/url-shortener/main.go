package main

import (
	"cthulhu/internal/config"
	"cthulhu/internal/lib/logger/sl"
	"cthulhu/internal/storage/sqlite"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// init config/ cleanenv (умеет читать с разных расширений. struct tag)
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)
	log.Info("starting cthulhu", slog.String("env", cfg.Env))

	storageCTX, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storageCTX

	// TODO: init storage/ sqlite

	// TODO: init router/ chi, "chi render"

	// TODO: run server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log

}
