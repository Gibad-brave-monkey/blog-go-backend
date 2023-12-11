package logger

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Gibad-brave-monkey/blog-go-backend/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

// NewLogger its a function-constructor

// Init initializes the logger (log/slog) based on the environment and log level.
func Init(env string, logLevel config.LogLevelStruct) (*slog.Logger, error) {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.Level(logLevel.LevelInfo),
		}))
	case envDev, envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.Level(logLevel.LevelInfo),
		}))
	default:
		return nil, fmt.Errorf("unknown environment: %s", env)
	}

	return log, nil
}
