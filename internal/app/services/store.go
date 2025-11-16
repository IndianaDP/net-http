package services

import (
	"github.com/IndianaDP/net-http/internal/app/db"
)

type URLStore interface {
	SaveURL(url string) (string, error)
	GetURL(uuid string) (string, error)
	GetAllUrls() ([]ResponseWriter, error)
}
type URLStoreService struct {
	conn db.DB
}

func NewURLStoreService(conn db.DB) *URLStoreService {
	return &URLStoreService{
		conn: conn,
	}
}
