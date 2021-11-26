package main

import (
	"hello-branch/greeting"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()

	router.Get("/api/v1/greeting", greeting.HelloText)

	router.Post("/api/v2/greeting", greeting.HelloJSON)

	http.ListenAndServe(":3000", router)
}
