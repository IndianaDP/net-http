package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IndianaDP/net-http/internal/app/utils"
)

func (h *Handlers) GetUrlsList(w http.ResponseWriter, r *http.Request) {
	urls, err := h.urlStore.GetURLs()
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, "Failed to get URLs list")
		return
	}

	if urls == nil {
		utils.ResponseError(w, http.StatusNotFound, "No URLs stored")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(urls)
}
