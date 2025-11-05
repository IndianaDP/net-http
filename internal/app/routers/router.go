package routers

import (
	"net/http"

	"github.com/IndianaDP/net-http/internal/app/config"
	"github.com/IndianaDP/net-http/internal/app/handlers"
	"github.com/IndianaDP/net-http/internal/app/services"
)

func SetupRouter(cfg *config.Values) http.Handler {
	store := services.NewURLStoreService()
	h := handlers.NewHandlers(cfg, store)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.URL.Path == "/" {
			h.CreateURLForRedirect(w, r)
			return
		}

		if r.Method == http.MethodGet {
			h.RedirectByID(w, r)
			return
		}

		http.Error(w, "Not found", http.StatusNotFound)
	})

	return mux
}
