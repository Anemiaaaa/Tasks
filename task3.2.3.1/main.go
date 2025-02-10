
package main

import (
"net/http"

"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.Get("/1", helloWorld)
	r.Get("/2", helloWorldTwo)
	r.Post("/3", helloWorldThree)

	http.ListenAndServe(":8080", r)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world"))
}

func helloWorldTwo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world 2"))
}

func helloWorldThree(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world 3"))
}