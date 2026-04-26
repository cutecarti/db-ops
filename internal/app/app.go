package app

import (
	"context"
	"log"

	"github.com/cutecarti/db-ops/internal/config"
	"github.com/cutecarti/db-ops/internal/db"
	router "github.com/cutecarti/db-ops/internal/delivery/http"
	"github.com/cutecarti/db-ops/internal/delivery/http/handlers"
	"github.com/cutecarti/db-ops/internal/repository"
	"github.com/cutecarti/db-ops/internal/server"
	"github.com/cutecarti/db-ops/internal/service"
)

func Run() {
	ctx := context.Background()
	cfg, err := config.GetCfg(ctx)
	if err != nil {
		log.Fatal("cant load config: ", err)
	}
	log.Println("init app")

	if err := db.RunMigrations(cfg.DB_URL); err != nil {
		log.Fatal(err)
	}

	pool, err := db.MakeNewPool(ctx, cfg.DB_URL)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer pool.Close()

	repo := repository.NewRecordRepo(pool)

	service := service.NewDBService(repo)

	recordHandler := handlers.NewRecordHandler(service)

	r := router.NewRouter(recordHandler)

	server := server.NewServer(":8000", r)
	log.Println("server started on :8000")

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
