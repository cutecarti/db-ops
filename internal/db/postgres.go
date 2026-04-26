package db

import (
	"context"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5/pgxpool"
)

func MakeNewPool(ctx context.Context, db_url string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, db_url)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}
	return pool, nil
}

func RunMigrations(databaseURL string) error {
	m, err := migrate.New(
		"file://internal/db/migrations",
		databaseURL,
	)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
