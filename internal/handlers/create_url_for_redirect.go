package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/eagle/net-http/internal/services"
)

const uuid = "i2o3hgo3ihg"

func CreateURLForRedirect(w http.ResponseWriter, r *http.Request) {
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

	services.SaveURL(uuid, string(url))

	fmt.Fprintf(w, "http://localhost:8080/%s", uuid)
}
