package routers

import (
	"net/http"

	"github.com/eagle/net-http/internal/handlers"
)

func SetupRouter() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.URL.Path == "/" {
			handlers.CreateURLForRedirect(w, r)
			return
		}

		if r.Method == http.MethodGet {
			handlers.RedirectByID(w, r)
			return
		}

		http.Error(w, "Not found", http.StatusNotFound)
	})
}
