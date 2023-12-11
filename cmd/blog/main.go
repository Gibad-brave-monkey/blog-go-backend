package main

import (
	"fmt"

	config "github.com/Gibad-brave-monkey/blog-go-backend/internal/config"
)

func main() {
	// config - cleanenv
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// logging - slog

	// router - chi

	// db - PostgresQL

	// run server
}
