package router

import (
	"net/http"

	"github.com/cutecarti/db-ops/internal/delivery/http/handlers"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/cutecarti/db-ops/docs"
)

func NewRouter(recordHandler *handlers.RecordHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /records", recordHandler.Create)

	mux.HandleFunc("GET /records/{id}", recordHandler.Get)

	mux.HandleFunc("DELETE /records/{id}", recordHandler.Delete)

	mux.HandleFunc("PUT /records/{id}", recordHandler.Update)

	mux.HandleFunc("GET /", recordHandler.Home)

	mux.HandleFunc("GET /docs/", httpSwagger.WrapHandler)

	return mux
}
