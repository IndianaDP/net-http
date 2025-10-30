package handlers

import (
	"github.com/IndianaDP/net-http/internal/app/config"
	"github.com/IndianaDP/net-http/internal/app/interfaces"
)

type Handlers struct {
	config   *config.Values
	urlStore interfaces.URLStore
}

func NewHandlers(cfg *config.Values, store interfaces.URLStore) *Handlers {
	h := &Handlers{
		config:   cfg,
		urlStore: store,
	}

	return h
}
