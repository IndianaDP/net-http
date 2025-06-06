package routers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRouter_PostAndGet(t *testing.T) {
	SetupRouter()

	postReq := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("https://example.com"))
	postW := httptest.NewRecorder()

	http.DefaultServeMux.ServeHTTP(postW, postReq)

	postResp := postW.Result()
	if postResp.StatusCode != http.StatusOK {
		t.Fatalf("POST: expected status 200, got %d", postResp.StatusCode)
	}

	postReq = httptest.NewRequest(http.MethodPost, "/someURL", strings.NewReader("https://example.com"))
	postW = httptest.NewRecorder()

	http.DefaultServeMux.ServeHTTP(postW, postReq)

	postResp = postW.Result()
	if postResp.StatusCode != http.StatusNotFound {
		t.Fatalf("POST: expected status 404, got %d", postResp.StatusCode)
	}

	getReq := httptest.NewRequest(http.MethodGet, "/i2o3hgo3ihg", nil)
	getW := httptest.NewRecorder()

	http.DefaultServeMux.ServeHTTP(getW, getReq)

	getResp := getW.Result()
	if getResp.StatusCode != http.StatusFound {
		t.Fatalf("GET: expected status 302, got %d", getResp.StatusCode)
	}
}
