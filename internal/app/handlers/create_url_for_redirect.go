package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (h *Handlers) CreateURLForRedirect(w http.ResponseWriter, r *http.Request) {
	hostname := h.config.Hostname
	uuid := h.config.UUID

	url, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read URL", http.StatusBadRequest)
		return
	}
	if len(url) == 0 {
		http.Error(w, "URL is empty", http.StatusBadRequest)
		return
	}
	if !strings.HasPrefix(string(url), "http://") && !strings.HasPrefix(string(url), "https://") {
		http.Error(w, "URL must start with http:// or https://", http.StatusBadRequest)
		return
	}

	h.urlStore.SaveURL(uuid, string(url))

	fmt.Fprintf(w, "%s/%s", hostname, uuid)
}
