package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func (h *Handlers) RedirectByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	url, err := h.urlStore.GetURL(id)
	if err != nil {
		http.Error(w, "Failed to get url for redirect", http.StatusInternalServerError)
		return
	}

	if url == "" {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	if !strings.HasPrefix(strings.TrimSpace(url), "https://") && !strings.HasPrefix(strings.TrimSpace(url), "http://") {
		http.Error(w, "Invalid URL stored", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Redirecting ID %s to URL: %s\n", id, url)

	http.Redirect(w, r, url, http.StatusFound)
}
