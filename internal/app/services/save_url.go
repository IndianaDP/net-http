package services

import (
	"crypto/sha1"
	"fmt"
)

func (s *URLStoreService) SaveURL(url string) (string, error) {
	uid := s.encode(url)

	id, err := s.conn.IsUrlExists(url)
	if err != nil {
		fmt.Printf("Error checking URL existence in DB: %v\n", err)
		return "", err
	}
	if id != "" {
		fmt.Printf("URL already exists in DB with UUID: %s\n", id)
		return "", nil
	}

	pk, err := s.conn.InsertURLIntoDB(url, uid)
	if err != nil {
		fmt.Printf("Error inserting URL into DB: %v\n", pk)
		return "", err
	}

	fmt.Printf("URL saved with id: %s\n", pk)

	count, err := s.conn.StoreCount()
	if err != nil {
		fmt.Printf("Error retrieving store count from DB: %v\n", err)
		return "", err
	}
	fmt.Printf("URLs count: %s \n", count)

	savedUrls, err := s.conn.StoredUrls()
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
