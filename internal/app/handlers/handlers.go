package handlers

import (
	"net/http"

	"github.com/IndianaDP/net-http/internal/app/config"
	"github.com/IndianaDP/net-http/internal/app/services"
)

type Handlers struct {
	config   *config.Values
	urlStore services.URLStore
}

func NewHandlers(cfg *config.Values, store services.URLStore) IHandlers {
	h := &Handlers{
		config:   cfg,
		urlStore: store,
	}

	return h
}

type IHandlers interface {
	CreateURLForRedirect(w http.ResponseWriter, r *http.Request)
	RedirectByID(w http.ResponseWriter, r *http.Request)
	GetUrlsList(w http.ResponseWriter, r *http.Request)
}
