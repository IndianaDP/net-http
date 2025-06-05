package services

import (
	"testing"
)

func TestSaveAndGetURL(t *testing.T) {
	id := "test-id"
	url := "https://example.com"

	SaveURL(id, url)

	saved, found := GetURL(id)
	if !found {
		t.Fatalf("Expected url to be found")
	}
	if saved != url {
		t.Errorf("Expected %s, got %s", url, saved)
	}
}
