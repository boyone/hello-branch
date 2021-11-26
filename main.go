package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()

	router.Get("/api/v1/greeting", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		fmt.Fprintf(w, "Hello %s", name)
	})

	http.ListenAndServe(":3000", router)
}
