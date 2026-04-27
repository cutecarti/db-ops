package app

import (
	"context"
	"log"

	"github.com/cutecarti/db-ops/internal/config"
	"github.com/cutecarti/db-ops/internal/repository/db"
)

func Run() {
	ctx := context.Background()

	cfg, err := config.GetCfg(ctx)
	if err != nil {
		log.Fatal("cant load config: ", err)
	}

	if err := db.RunMigrations(cfg.DB_URL); err != nil {
		log.Fatal(err)
	}

	srv, cleanup, err := InitServer(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	log.Println("server started on :8000")

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
