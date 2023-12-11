package main

import (
	"fmt"

	config "github.com/Gibad-brave-monkey/blog-go-backend/internal/config"
	"github.com/Gibad-brave-monkey/blog-go-backend/pkg/logger"
)

func main() {
	// config
	cfg := config.MustLoad()
	fmt.Println(cfg)
	// logger
	log, err := logger.Init(cfg.Env, cfg.LogLevel)
	if err != nil {
		fmt.Println(err)
	}

	log.Info("Info")
	log.Warn("warn")
	log.Error("error")

	// router - chi

	// db - PostgresQL

	// run server
}
