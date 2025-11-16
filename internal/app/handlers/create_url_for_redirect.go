package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (h *Handlers) CreateURLForRedirect(w http.ResponseWriter, r *http.Request) {
	hostname := h.config.Hostname

	url, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read URL", http.StatusBadRequest)
		return
	}
	if len(url) == 0 {
		http.Error(w, "URL is empty", http.StatusBadRequest)
		return
	}

	urlString := string(url)

	if !strings.HasPrefix(urlString, "http://") && !strings.HasPrefix(urlString, "https://") {
		http.Error(w, "URL must start with http:// or https://", http.StatusBadRequest)
		return
	}

	uid, err := h.urlStore.SaveURL(urlString)
	if err != nil {
		http.Error(w, "Failed to save URL", http.StatusInternalServerError)
		return
	}
	if uid == "" {
		http.Error(w, "URL already exists", http.StatusConflict)
		return
	}

	fmt.Fprintf(w, "%s/%s", hostname, uid)
}
