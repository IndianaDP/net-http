package services

type URLStoreService struct {
	store map[string]string
}

func NewURLStoreService() *URLStoreService {
	return &URLStoreService{
		store: make(map[string]string),
	}
}

func (s *URLStoreService) SaveURL(uuid, url string) {
	s.store[uuid] = url
}

func (s *URLStoreService) GetURL(uuid string) (string, bool) {
	url, ok := s.store[uuid]
	return url, ok
}

func (s *URLStoreService) IsStoreEmpty() bool {
	return len(s.store) == 0
}
