package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/IndianaDP/net-http/internal/config"
	"github.com/IndianaDP/net-http/internal/services"
)

func CreateURLForRedirect(w http.ResponseWriter, r *http.Request) {

	cfg, err := config.LoadConfig(false)
	if err != nil {
		http.Error(w, "Failed to load config", http.StatusInternalServerError)
		return
	}
	hostname := string(cfg.Hostname)
	uuid := string(cfg.UUID)

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

	fmt.Fprintf(w, "%s/%s", hostname, uuid)
}
