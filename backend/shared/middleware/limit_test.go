package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestLimiterRejectsAfterConfiguredMaximum(t *testing.T) {
	limiter := NewLimiter()
	handler := limiter.Middleware("test", 2, time.Minute, func(*http.Request) string { return "same" }, nil)(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusNoContent) }),
	)
	for _, expected := range []int{http.StatusNoContent, http.StatusNoContent, http.StatusTooManyRequests} {
		response := httptest.NewRecorder()
		handler.ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/", nil))
		if response.Code != expected {
			t.Fatalf("status = %d, want %d", response.Code, expected)
		}
	}
}

func TestDisabledLimiterDoesNotCount(t *testing.T) {
	handler := NewLimiter().Middleware("off", 0, time.Minute, IPKey, nil)(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusNoContent) }),
	)
	for range 3 {
		response := httptest.NewRecorder()
		handler.ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/", nil))
		if response.Code != http.StatusNoContent {
			t.Fatalf("status = %d", response.Code)
		}
	}
}
