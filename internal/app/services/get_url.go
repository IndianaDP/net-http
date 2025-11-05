package services

func (s *URLStoreService) GetURL(uuid string) (string, bool) {
	url, ok := s.store[uuid]
	return url, ok
}
