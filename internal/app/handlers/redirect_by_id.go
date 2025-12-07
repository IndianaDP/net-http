package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/IndianaDP/net-http/internal/app/utils"
	"github.com/go-chi/chi/v5"
)

func (h *Handlers) RedirectByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		utils.ResponseError(w, http.StatusBadRequest, "ID is required")
		return
	}

	url, err := h.urlStore.GetURL(id)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, "Failed to get url for redirect")
		return
	}

	if url == "" {
		utils.ResponseError(w, http.StatusNotFound, "URL not found")
		return
	}

	if !strings.HasPrefix(strings.TrimSpace(url), "https://") && !strings.HasPrefix(strings.TrimSpace(url), "http://") {
		utils.ResponseError(w, http.StatusInternalServerError, "Invalid URL stored")
		return
	}

	fmt.Printf("Redirecting ID %s to URL: %s\n", id, url)

	http.Redirect(w, r, url, http.StatusFound)
}
