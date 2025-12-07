package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/IndianaDP/net-http/internal/app/config"
	"github.com/IndianaDP/net-http/internal/app/db/storage"
	"github.com/IndianaDP/net-http/internal/app/handlers"
	"github.com/IndianaDP/net-http/internal/app/services"
)

func SetupRouter(cfg *config.Values, storage storage.Storage) http.Handler {
	store := services.NewURLStoreService(storage)
	h := handlers.NewHandlers(cfg, store)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Post("/create", h.CreateURLForRedirect)
	router.Get("/list", h.GetUrlsList)
	router.Get("/", h.RedirectByID)
	router.Get("/{id}", h.RedirectByID)

	return router
}
