package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handlers) GetUrlsList(w http.ResponseWriter, r *http.Request) {
	urls, err := h.urlStore.GetURLs()
	if err != nil {
		http.Error(w, "Failed to get URLs list", http.StatusInternalServerError)
		return
	}

	if urls == nil {
		http.Error(w, "No URLs stored", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(urls)
}
