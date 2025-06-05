package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateURLForRedirect(t *testing.T) {

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("https://example.com"))
	w := httptest.NewRecorder()

	CreateURLForRedirect(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Result().StatusCode)
	}

	if !strings.Contains(w.Body.String(), "http://localhost:8080/i2o3hgo3ihg") {
		t.Errorf("Expected response to contain redirect URL, got %s", w.Body.String())
	}

	req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
	w = httptest.NewRecorder()

	CreateURLForRedirect(w, req)

	if w.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("Expected 400 for empty URL, got %d", w.Result().StatusCode)
	}

	req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("example.com"))
	w = httptest.NewRecorder()
	CreateURLForRedirect(w, req)
	if w.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("Expected 400 for URL without https:// or http://, got %d", w.Result().StatusCode)
	}
}
