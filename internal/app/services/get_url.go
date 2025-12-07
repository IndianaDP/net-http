package services

import (
	"fmt"
)

func (s *URLStoreService) GetURL(uuid string) (string, error) {

	url, err := s.storage.Get(uuid)
	if err != nil {
		fmt.Printf("Error retrieving URL from DB: %v\n", err)
		return "", err
	}

	if url == "" {
		fmt.Printf("No URL found for UUID: %s\n", uuid)
		return "", nil
	}

	return url, nil
}
