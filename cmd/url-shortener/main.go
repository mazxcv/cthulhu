package main

import (
	"cthulhu/internal/config"
	mwLogger "cthulhu/internal/http-server/middleware/logger"
	"cthulhu/internal/lib/logger/handlers/slogpretty"
	"cthulhu/internal/lib/logger/sl"
	"cthulhu/internal/storage/sqlite"
	"log/slog"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	log.Info("Information starting cthulhu", slog.String("env", cfg.Env))
	log.Debug("Debug starting cthulhu", slog.String("env", cfg.Env))
	log.Warn("Warning starting cthulhu", slog.String("env", cfg.Env))
	log.Error("Error starting cthulhu", slog.String("env", cfg.Env))

	storageCTX, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storageCTX

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	// TODO: run server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = setupPrettySlog()
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log

}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)
	return slog.New(handler)
}
