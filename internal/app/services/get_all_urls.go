package services

import (
	"fmt"
)

type ResponseWriter struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

func (s *URLStoreService) GetAllUrls() ([]ResponseWriter, error) {

	urls, err := s.conn.StoredUrls()
	if err != nil {
		fmt.Printf("Error retrieving URLs from DB: %v\n", err)
		return nil, err
	}

	if urls == nil {
		fmt.Printf("No URLs found in the DB\n")
		return nil, nil
	}

	var response []ResponseWriter
	for _, u := range urls {
		response = append(response, ResponseWriter{
			ID:  u.UUID,
			URL: u.URL,
		})
	}

	return response, nil
}
