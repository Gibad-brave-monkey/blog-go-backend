package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	config "github.com/Gibad-brave-monkey/blog-go-backend/internal/config"
	"github.com/Gibad-brave-monkey/blog-go-backend/internal/db/postgresdb"
	"github.com/Gibad-brave-monkey/blog-go-backend/internal/routers"
	"github.com/Gibad-brave-monkey/blog-go-backend/pkg/logger"
	"github.com/go-chi/chi/v5"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 9*time.Second)
	defer cancel()
	// config
	cfg := config.MustLoad()
	fmt.Println(cfg)
	// logger
	log, err := logger.Init(cfg.Env, cfg.LogLevel)
	if err != nil {
		fmt.Println(err)
	}

	db, err := postgresdb.NewPostgresDB(ctx, log, cfg.DB)
	if err != nil {
		log.Error("Failed to connect db", err)
	}

	_ = db

	router := chi.NewRouter()

	routers.BindRoutesFn(log, router)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%v", cfg.Port),
		Handler:      router,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	// Run HTTP-server
	log.Info("Server is running on the:", slog.Int("port:", cfg.Port))
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
