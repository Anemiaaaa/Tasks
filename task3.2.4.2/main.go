package main

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func main() {
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

	http.ListenAndServe(":8080", r)
}

func LoggerMiddleware(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			logger.Info("request",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.String("ip", r.RemoteAddr),
				zap.Duration("duration", time.Since(start)),
			)
			next.ServeHTTP(w, r)
		})
	}
}
