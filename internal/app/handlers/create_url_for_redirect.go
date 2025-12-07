package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/IndianaDP/net-http/internal/app/utils"
)

type RequestBody struct {
	URL string `json:"url"`
}

func (h *Handlers) CreateURLForRedirect(w http.ResponseWriter, r *http.Request) {
	hostname := h.config.Hostname

	jsonBody := &RequestBody{}
	err := json.NewDecoder(r.Body).Decode(jsonBody)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	if len(jsonBody.URL) == 0 {
		http.Error(w, "URL field is missing", http.StatusBadRequest)
		return
	}

	urlString := string([]byte(jsonBody.URL))

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
