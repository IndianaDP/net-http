package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/IndianaDP/net-http/internal/app/utils"
)

type RequestBody struct {
	UserID string `json:"user_id"`
	URL    string `json:"url"`
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
		utils.ResponseError(w, http.StatusBadRequest, "URL field is missing")
		return
	}

	urlString := string([]byte(jsonBody.URL))
	userIdString := string([]byte(jsonBody.UserID))

	if !strings.HasPrefix(urlString, "http://") && !strings.HasPrefix(urlString, "https://") {
		utils.ResponseError(w, http.StatusBadRequest, "URL must start with http:// or https://")
		return
	}

	uid, err := h.urlStore.SaveURL(userIdString, urlString)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, "Failed to save URL")
		return
	}
	if uid == "" {
		utils.ResponseError(w, http.StatusConflict, "URL already exists")
		return
	}

	fmt.Fprintf(w, "%s/%s", hostname, uid)
}
