package handlers

import (
	"net/http"
	"strings"

	"github.com/IndianaDP/net-http/internal/services"
)

func RedirectByID(w http.ResponseWriter, r *http.Request) {

	if services.IsStoreEmpty() {
		http.Error(w, "URL database is empty", http.StatusServiceUnavailable)
		return
	}

	id := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/"))

	if r.URL.Path == "/" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	url, ok := services.GetURL(id)
	if !ok {
		http.Error(w, "Invalid ID", http.StatusNotFound)
		return
	}

	if strings.TrimSpace(url) == "" {
		http.Error(w, "Invalid URL stored", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}
