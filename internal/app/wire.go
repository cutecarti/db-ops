//go:build wireinject
// +build wireinject

package app

import (
	"context"
	"net/http"

	"github.com/cutecarti/db-ops/internal/config"
	router "github.com/cutecarti/db-ops/internal/delivery/http"
	"github.com/cutecarti/db-ops/internal/delivery/http/handlers"
	"github.com/cutecarti/db-ops/internal/repository"
	"github.com/cutecarti/db-ops/internal/repository/db"
	"github.com/cutecarti/db-ops/internal/server"
	"github.com/cutecarti/db-ops/internal/service"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitServer(ctx context.Context) (*server.Server, func(), error) {
	panic(wire.Build(
		config.GetCfg,
		providePool,

		repository.NewRecordRepo,
		wire.Bind(new(service.RecordRepository), new(*repository.RecordRepo)),

		service.NewDBService,
		handlers.NewRecordHandler,
		router.NewRouter,
		provideServer,
	))
}

func providePool(ctx context.Context, cfg *config.Settings) (*pgxpool.Pool, func(), error) {
	pool, err := db.MakeNewPool(ctx, cfg.DB_URL)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		pool.Close()
	}

	return pool, cleanup, nil
}

func provideServer(r http.Handler) *server.Server {
	return server.NewServer(":8000", r)
}
