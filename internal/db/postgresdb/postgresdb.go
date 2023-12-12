package postgresdb

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/Gibad-brave-monkey/blog-go-backend/internal/config"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewPostgresDB(ctx context.Context, logger *slog.Logger, cfg config.DBConfig) (*PostgresDB, error) {
	const op = "db.postgresdb.NewPostgresDB"

	var connectString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", connectString)
	if err != nil {
		logger.Error("Failed to Connecting DB", err)
		return &PostgresDB{}, fmt.Errorf("%v:%w", op, err)
	}

	if err := db.Ping(); err != nil {
		logger.Error("Failed to PING DB", err)
		return &PostgresDB{}, fmt.Errorf("%v: %w", op, err)
	}

	defer db.Close()

	logger.Info("Connected to DB!")

	return &PostgresDB{db: db, logger: logger}, nil
}

// Close DB connection
func (p *PostgresDB) Close() {
	if p.db != nil {
		p.db.Close()
		p.logger.Info("Closed DB connection")
	}
}
