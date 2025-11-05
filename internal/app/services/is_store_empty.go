package services

func (s *URLStoreService) IsStoreEmpty() bool {
	return len(s.store) == 0
}
