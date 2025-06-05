package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IndianaDP/net-http/internal/services"
)

func TestRedirectByID(t *testing.T) {
	services.SaveURL("i2o3hgo3ihg", "https://example.com")

	req := httptest.NewRequest(http.MethodGet, "/i2o3hgo3ihg", nil)
	w := httptest.NewRecorder()

	RedirectByID(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusFound {
		t.Errorf("Expected 302, got %d", resp.StatusCode)
	}

	redirectLocation := resp.Header.Get("Location")
	if redirectLocation != "https://example.com" {
		t.Errorf("Expected location is 'https://example.com', got '%s'", redirectLocation)
	}

	req = httptest.NewRequest(http.MethodGet, "/123", nil)
	w = httptest.NewRecorder()

	RedirectByID(w, req)

	if w.Result().StatusCode != http.StatusNotFound {
		t.Errorf("expected 404 because ID does not exist, got %d", w.Result().StatusCode)
	}
	req = httptest.NewRequest(http.MethodGet, "/", nil)
	w = httptest.NewRecorder()

	RedirectByID(w, req)

	if w.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400 for empty ID, got %d", w.Result().StatusCode)
	}

	services.SaveURL("i2o3hgo3ihg", "")

	req = httptest.NewRequest(http.MethodGet, "/i2o3hgo3ihg", nil)
	w = httptest.NewRecorder()

	RedirectByID(w, req)

	if w.Result().StatusCode != http.StatusInternalServerError {
		t.Errorf("expected 404 for empty URL, got %d", w.Result().StatusCode)
	}
}
