package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestRoutes(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	r := chi.NewRouter()
	r.Use(LoggerMiddleware(logger))

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	r.Get("/amir", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Amir"))
	})
	r.Post("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Data posted"))
	})

	tests := []struct {
		name         string
		method       string
		path         string
		expectedCode int
		expectedBody string
	}{
		{"Route /hello", "GET", "/hello", http.StatusOK, "Hello, World!"},
		{"Route /amir", "GET", "/amir", http.StatusOK, "Amir"},
		{"Route /data", "POST", "/data", http.StatusOK, "Data posted"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.Equal(t, tt.expectedBody, w.Body.String())
		})
	}
}