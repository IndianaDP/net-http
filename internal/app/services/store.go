package services

type URLStore interface {
	SaveURL(url string) string
	GetURL(uuid string) (string, bool)
	IsStoreEmpty() bool
}
type URLStoreService struct {
	store map[string]string
}

func NewURLStoreService() *URLStoreService {
	return &URLStoreService{
		store: make(map[string]string),
	}
}
