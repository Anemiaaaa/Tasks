package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutes(t *testing.T) {
	r := chi.NewRouter()
	r.Get("/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello world"))
	})

	r.Get("/2", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello world 2"))
	})

	r.Get("/3", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello world 3"))
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
	})

	tests := []struct {
		name         string
		path         string
		expectedCode int
		expectedBody string
	}{
		{"Route /1", "/1", http.StatusOK, "Hello world"},
		{"Route /2", "/2", http.StatusOK, "Hello world 2"},
		{"Route /3", "/3", http.StatusOK, "Hello world 3"},
		{"Route / (not found)", "/", http.StatusNotFound, "Not Found"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", tt.path, nil)
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