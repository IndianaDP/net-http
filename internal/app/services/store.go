package services

import (
	"github.com/IndianaDP/net-http/internal/app/db/storage"
)

type URLStore interface {
	SaveURL(user_id string, url string) (string, error)
	GetURL(uuid string) (string, error)
	GetURLs() ([]ResponseWriter, error)
}
type URLStoreService struct {
	storage storage.Storage
}

func NewURLStoreService(storage storage.Storage) *URLStoreService {
	return &URLStoreService{
		storage: storage,
	}
}
