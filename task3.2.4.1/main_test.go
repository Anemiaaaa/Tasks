package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutes(t *testing.T) {
	r := chi.NewRouter()
	r.Use(LoggerMiddleware)
	r.Get("/route1", handleRoute1)
	r.Post("/route2", handleRoute2)
	r.Put("/route3", handleRoute3)

	tests := []struct {
		method       string
		path         string
		expectedCode int
		expectedBody string
	}{
		{"GET", "/route1", http.StatusOK, "Route 1 - GET request"},
		{"POST", "/route2", http.StatusOK, "Route 2 - POST request"},
		{"PUT", "/route3", http.StatusOK, "Route 3 - PUT request"},
	}

	for _, tt := range tests {
		t.Run(tt.method+" "+tt.path, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code != tt.expectedCode {
				t.Errorf("expected %d, got %d", tt.expectedCode, w.Code)
			}

			if w.Body.String() != tt.expectedBody {
				t.Errorf("expected body %q, got %q", tt.expectedBody, w.Body.String())
			}
		})
	}
}