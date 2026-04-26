package main

import (
	"log"

	"github.com/cutecarti/db-ops/internal/app"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}
}

// @title DB Ops API
// @version 1.0
// @description API for records.
// @host localhost:8000
// @BasePath /
func main() {
	app.Run()
}
