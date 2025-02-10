package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(LoggerMiddleware)

	r.Get("/route1", handleRoute1)
	r.Post("/route2", handleRoute2)
	r.Put("/route3", handleRoute3)

	log.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", r)
}

func handleRoute1(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Route 1 - GET request"))
}

func handleRoute2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Route 2 - POST request"))
}

func handleRoute3(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Route 3 - PUT request"))
}
