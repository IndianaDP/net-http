package services

import (
	"crypto/sha256"
	"fmt"
)

func (s *URLStoreService) SaveURL(url string) string {
	uid := s.encode(url)

	if _, ok := s.store[uid]; ok {
		fmt.Printf("URL: %s already exists \n", url)
		return uid
	}

	s.store[uid] = url

	fmt.Printf("Store count: %d \n", len(s.store))

	for k, v := range s.store {
		fmt.Printf("Saved UUID: %s with URL: %s\n", k, v)
	}

	return uid
}

func (s *URLStoreService) encode(value string) string {
	hash := sha256.Sum256([]byte(value))

	return fmt.Sprintf("%x", hash)[:10]
}
