package config

import (
	"context"
	"fmt"
	"os"
)

type Settings struct {
	DB_URL string
}

func GetCfg(ctx context.Context) (*Settings, error) {
	user, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		return nil, fmt.Errorf("missing POSTGRES_USER var")
	}
	password, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		return nil, fmt.Errorf("missing POSTGRES_PASSWORD var")
	}
	db_name, ok := os.LookupEnv("POSTGRES_DB")
	if !ok {
		return nil, fmt.Errorf("missing POSTGRES_DB var")
	}
	port, ok := os.LookupEnv("POSTGRES_PORT")
	if !ok {
		return nil, fmt.Errorf("missing POSTGRES_PORT var")
	}

	return &Settings{
		DB_URL: "postgres://" + user + ":" + password + "@" + "db:" + port + "/" + db_name + "?sslmode=disable",
	}, nil
}
