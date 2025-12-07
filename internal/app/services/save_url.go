package services

import (
	"crypto/sha1"
	"fmt"
)

func (s *URLStoreService) SaveURL(url string) (string, error) {
	uid := s.encode(url)

	exists, err := s.storage.Exists(uid)
	if err != nil {
		fmt.Printf("Error checking URL existence in DB: %v\n", err)
		return "", err
	}
	if exists {
		fmt.Printf("URL already exists in DB")
		return "", nil
	}

	err = s.storage.Insert(url, uid)
	if err != nil {
		return "", err
	}

	count, err := s.storage.Count()
	if err != nil {
		fmt.Printf("Error retrieving store count from DB: %v\n", err)
		return "", err
	}
	fmt.Printf("URLs count: %s \n", count)

	savedUrls, err := s.storage.List()
	if err != nil {
		fmt.Printf("Error retrieving stored URLs from DB: %v\n", err)
		return "", err
	}

	for _, v := range savedUrls {
		fmt.Printf("Saved UUID: %s with URL: %s\n", v.UUID, v.URL)
	}

	return uid, nil
}

func (s *URLStoreService) encode(value string) string {
	hash := sha1.Sum([]byte(value))

	return fmt.Sprintf("%x", hash)
}
