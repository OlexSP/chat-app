package database

import (
	"backend/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

// ConnectPostgres establishes a connection to the PostgresQL database.
func ConnectPostgres(cfg *config.PostgresConfig, maxAttempts int) (pool *pgxpool.Pool, err error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
	)

	err = doWithTries(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, connStr)
		if err != nil {
			return err
		}
		return nil

	}, maxAttempts, time.Second)

	if err != nil {
		return nil, fmt.Errorf("unable to connect to PostgreSQL: %w", err)
	}

	return pool, nil
}

func doWithTries(fn func() error, tries int, delay time.Duration) (err error) {
	for tries > 0 {
		if err := fn(); err != nil {
			time.Sleep(delay)
			tries--

			continue
		}
		return nil
	}
	return
}
