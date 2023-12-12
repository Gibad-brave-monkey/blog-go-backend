package routers

import (
	"log/slog"

	recordsHandlers "github.com/Gibad-brave-monkey/blog-go-backend/internal/handlers/records-handlers"
	"github.com/go-chi/chi/v5"
)

func BindRoutesFn(log *slog.Logger, router *chi.Mux) {
	rh := recordsHandlers.RecordsHandlers{}
	// Logging GAVNA
	log.Info("Binds Routers")

	router.Post("/v1/save", rh.SaveRecord)
	router.Get("/v1/records", rh.GetRecords)
}
